package main

import (
	"github.com/labstack/echo/v4"
	"github.com/georgij1/TestTakGoBackendMedods/internal/routes"
	"github.com/georgij1/TestTakGoBackendMedods/pkg/middleware"
)

// @title My Echo API
// @version 1.0
func main() {
	e := echo.New()
	
	// Middleware
	e.Use(middleware.Logger())
	
	// Routes
	routes.RegisterAPI(e)
	
	e.Logger.Fatal(e.Start(":8080"))
}	