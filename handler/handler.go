package handler

import (
	"myapp/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func New(db *gorm.DB) Handler {
	return Handler{db}
}

func (handler Handler) GetTodos(c echo.Context) error {
	var todos []models.Todo
	tx := handler.db.Find(&todos)
	if tx.Error != nil {
		c.Logger().Error(tx.Error)
		return c.String(500, "error in GET /todos/")
	}
	return c.JSON(200, todos)
}
