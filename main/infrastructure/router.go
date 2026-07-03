package infrastructure

import (
	"notes/main/interfeces/api"
	"os"

	"github.com/labstack/echo/v4"
)

func Init() {
	e := echo.New()
	sqlHandler := NewSqlHandler()
	noteController := api.NewNoteController(sqlHandler)

	e.POST("/notes", noteController.CreateNode)
	e.GET("/notes", noteController.FindAll)
	e.DELETE("/notes/:id", noteController.DeleteById)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
