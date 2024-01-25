package middleware

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

func createLogger(c echo.Context, v middleware.RequestLoggerValues) error {
	l := zerolog.New(os.Stdout)
	l.Info().
		Str("URI", v.URI).
		Int("Status", v.Status).
		Msg("request")

	return nil
}

func requestLogger() echo.MiddlewareFunc {
	lc := middleware.RequestLoggerConfig{
		LogURI:        true,
		LogStatus:     true,
		LogValuesFunc: createLogger,
	}

	return middleware.RequestLoggerWithConfig(lc)
}
