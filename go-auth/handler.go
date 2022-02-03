package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type handler struct{}

// Most of the code is taken from the echo guide
// https://echo.labstack.com/cookbook/jwt
func (h *handler) login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "kurniadi" && password == "1234" {
		token := jwt.New(jwt.SigningMethodHS256)

		sessions := token.Claims.(jwt.MapClaims)
		sessions["name"] = "Jon Doe"
		sessions["admin"] = true
		sessions["exp"] = time.Now().Add(time.Hour * 72).Unix()

		t, err := token.SignedString([]byte("secret"))

		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}
	return echo.ErrUnauthorized
}

func (h *handler) profile(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	sessions := user.Claims.(jwt.MapClaims)

	name := sessions["name"].(string)
	exp := sessions["exp"].(float64)

	return c.JSON(http.StatusOK, map[string]string{
		"name":      name,
		"timestamp": fmt.Sprintf("%v", exp),
	})
}
