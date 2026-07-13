package infrastructure

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"notes/main/interfeces/api"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() {
	e := echo.New()
	sqlHandler := NewSqlHandler()
	if err := RunMigration(sqlHandler); err != nil {
		e.Logger.Fatalf("Failed to run migrations: %v", err)
	}
	noteController := api.NewNoteController(sqlHandler)

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogMethod:  true,
		LogURI:     true,
		LogStatus:  true,
		LogLatency: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			fmt.Printf("\x1b[32m[API REQUEST]\x1b[0m %s %s | Статус: %d | Время: %v\n",
				v.Method, v.URI, v.Status, v.Latency)
			return nil
		},
	}))

	e.Use(middleware.Recover())

	e.POST("/notes", noteController.CreateNode)
	e.GET("/notes", noteController.FindAll)
	e.DELETE("/notes/:id", noteController.DeleteById)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	go func() {
		if err := e.Start(":" + port); err != nil && !errors.Is(err, http.ErrServerClosed) {
			e.Logger.Fatalf("shutting down the server: %v", err)
		}
	}()
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	e.Logger.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatalf("could not shutdown server gracefully: %v", err)
	} else {
		e.Logger.Info("server shutdown gracefully")
	}
}
