package main

import (
	"github.com/labstack/echo/v4"
	"github.com/padulkemid/nusibes/internal/app"
	"github.com/padulkemid/nusibes/pkg/middleware"
)

const PORT = "localhost:9119"

func serverStart() {
	// Add instance
	e := echo.New()

	// Register middleware
	middleware.RegisterMiddleware(e)

	// Register routes
	app.RegisterRoutes(e)

	// Instance settings
	e.HideBanner = true

	// Start server
	e.Logger.Fatal(e.Start(PORT))
}
