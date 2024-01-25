package healthcheck

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const PING = "/ping"

func healthCheck(c echo.Context) error {
	pong := map[string]interface{}{
		"data": "pong!",
	}

	return c.JSON(http.StatusOK, pong)
}

func Routes(e *echo.Echo) {
	e.GET(PING, healthCheck)
}
