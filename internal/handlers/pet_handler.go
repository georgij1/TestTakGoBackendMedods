package handlers

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
	"github.com/yourusername/my-echo-app/internal/models"
)

// GetPetByID godoc
// @Summary Get pet by ID
// @Description Get pet details
// @Tags pets
// @Produce json
// @Param id path int true "Pet ID"
// @Success 200 {object} models.Pet
// @Router /api/pets/{id} [get]
func GetPetByID(c echo.Context) error {
	// Логика обработчика
	return c.JSON(http.StatusOK, models.Pet{ID: 1, Name: "Fluffy"})
}