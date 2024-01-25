package app

import (
	"github.com/labstack/echo/v4"
	"github.com/padulkemid/nusibes/internal/app/healthcheck"
)

func RegisterRoutes(e *echo.Echo) {
	healthcheck.Routes(e)
}
