package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleGet(c echo.Context) error {
	return c.String(http.StatusOK, "Example! GET")
}

func HandlePost(c echo.Context) error {
	return c.String(http.StatusCreated, "Example! POST")
}

func HandlePut(c echo.Context) error {
	return c.String(http.StatusOK, "Example! PUT")
}

func HandleDelete(c echo.Context) error {
	return c.String(http.StatusNoContent, "Example! DELETE")
}
