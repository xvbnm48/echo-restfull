package controllers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/xvbnm48/echo-restfull/helpers"
	"github.com/xvbnm48/echo-restfull/models"
)

type login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GenerateHashPassword(c echo.Context) error {
	// password := &login{}
	// if err := c.Bind(password); err != nil {
	// 	return c.JSON(http.StatusBadRequest, map[string]string{
	// 		"msg": err.Error(),
	// 	})
	// }

	password := c.Param("password")
	hash, _ := helpers.HashPassword(password)
	return c.JSON(http.StatusOK, hash)

}

func CheckLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	res, err := models.CheckLogin(username, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"msg": err.Error(),
		})
	}

	if !res {
		return echo.ErrUnauthorized
	}

	// token generate
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["level"] = "application"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// generate encoded token and send it as response.
	t, err := token.SignedString([]byte("sakuraaaaaaaaaaaaaaaaaaa"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"msg":   "Login Success",
		"token": t,
	})

	// return c.JSON(http.StatusOK, map[string]string{
	// 	"msg": "Login Success",
	// })
}
