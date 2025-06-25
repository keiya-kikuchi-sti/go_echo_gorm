package main

import (
	"myapp/handler"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetTodos(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/todos/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	db, err := run()
	assert.NoError(t, err)
	hd := handler.New(db)
	err = hd.GetTodos(c)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
