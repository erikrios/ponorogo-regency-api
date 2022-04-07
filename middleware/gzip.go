package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Gzip(e *echo.Echo) {
	e.Use(middleware.Gzip())
}
