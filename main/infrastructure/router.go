package infrastructure

import (
	"notes/main/interfeces/api"

	"github.com/labstack/echo/v4"
)

func Init() {
	e := echo.New()
	sqlHandler := NewSqlHandler()
	noteController := api.NewNoteController(sqlHandler)

	e.POST("/notes", noteController.CreateNode)
	e.GET("/notes", noteController.FindAll)
	e.DELETE("/notes", noteController.DeleteById)

	e.Logger.Fatal(e.Start(":8080"))
}
