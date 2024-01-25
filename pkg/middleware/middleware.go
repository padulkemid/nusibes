package middleware

import (
	"github.com/labstack/echo/v4"
)

func RegisterMiddleware(e *echo.Echo) {
	e.Use(requestLogger())
}
