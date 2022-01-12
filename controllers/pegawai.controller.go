package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/xvbnm48/echo-restfull/models"
)

func FetchAllPegawai(c echo.Context) error {
	result, err := models.FetchPegawai()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
