package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type handler struct{}

func (h *handler) login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "kurniadi" && password == "1234" {
		token := jwt.New(jwt.SigningMethodHS256)

		t, err := token.SignedString([]byte("secret"))

		session := token.Claims.(jwt.MapClaims)
		session["name"] = "Kurniadi Ahmad Wijaya"
		session["admin"] = true
		session["expired"] = time.Now().Add(time.Hour * 72).Unix()
		session["token"] = t

		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, session)
	}

	return echo.ErrUnauthorized
}
