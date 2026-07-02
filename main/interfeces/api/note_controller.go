package api

import (
	"net/http"
	"notes/main/domain"
	"notes/main/usecase"

	"github.com/labstack/echo/v4"
)

type NoteController struct {
	Interactor usecase.NoteInteractor
}

func NewNoteController(repo usecase.NoteRepository) *NoteController {
	return &NoteController{Interactor: usecase.NoteInteractor{NoteRepository: repo}}

}
func (controller *NoteController) CreateNode(c echo.Context) error {
	var note domain.Note
	if err := c.Bind(&note); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := controller.Interactor.Add(&note); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, note)
}

func (controller *NoteController) FindAll(c echo.Context) error {
	notes, err := controller.Interactor.FindAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, notes)
}

func (controller *NoteController) DeleteById(c echo.Context) error {
	id := c.Param("id")
	if err := controller.Interactor.DeleteById(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusNoContent, nil)
}
