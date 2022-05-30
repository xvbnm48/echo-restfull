package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/xvbnm48/echo-restfull/controllers"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "Hello World"})
	})
	e.GET("/pegawai", controllers.FetchAllPegawai)
	e.POST("/pegawai", controllers.StorePegawai)
	e.PUT("pegawai/:id", controllers.UpdatePegawai)
	e.DELETE("pegawai/:id", controllers.DeletePegawai)
	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	e.POST("/login", controllers.CheckLogin)
	return e
}
