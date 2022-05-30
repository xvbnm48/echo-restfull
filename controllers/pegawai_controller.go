package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/xvbnm48/echo-restfull/models"
)

func FetchAllPegawai(c echo.Context) error {
	result, err := models.FetchAllPegawai()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"msg": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func StorePegawai(c echo.Context) error {
	// nama := c.FormValue("nama")
	// alamat := c.FormValue("alamat")
	// telepon := c.FormValue("telepon")
	p := &models.Pegawai{}
	if err := c.Bind(p); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"msg": err.Error(),
		})
	}

	result, err := models.StorePegawai(p.Nama, p.Alamat, p.Telepon)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"msg": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdatePegawai(c echo.Context) error {
	id := c.Param("id")
	p := &models.Pegawai{}
	if err := c.Bind(p); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"msg": err.Error(),
		})
	}
	conv_id, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"msg": err.Error(),
		})
	}

	result, err := models.UpdatePegawai(conv_id, p.Nama, p.Alamat, p.Telepon)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"msg": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func DeletePegawai(c echo.Context) error {
	id := c.Param("id")
	p := &models.Pegawai{}
	if err := c.Bind(p); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"msg": err.Error(),
		})
	}

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"msg": err.Error(),
		})
	}

	result, err := models.DeletePegawai(conv_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"msg": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}
