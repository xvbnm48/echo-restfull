package routes

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/xvbnm48/echo-restfull/controllers"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/pegawai", controllers.FetchAllPegawai)

	return e
}
