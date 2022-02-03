package main

import (
	"github.com/labstack/echo/middleware"
)

var IsLogged = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("secret"),
})
