package routes

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/yourusername/my-echo-app/internal/handlers"
)

func RegisterAPI(e *echo.Echo) {
	// Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	
	// API Routes
	api := e.Group("/api")
	{
		pets := api.Group("/pets")
		{
			pets.GET("/:id", handlers.GetPetByID)
		}
	}
}