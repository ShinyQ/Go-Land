package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	h := &handler{}

	e.POST("/login", h.login)
	e.GET("/home", h.profile, IsLogged)
	e.GET("/admin", h.profile, IsLogged, isAdmin)

	e.Logger.Fatal(e.Start(":1323"))
}
