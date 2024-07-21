package handlers

import (
	"github.com/gofiber/fiber/v2"
	"namaz-vakitleri/models"
	"namaz-vakitleri/services"
)

type CityHandler struct {
	Service *services.CityService
}

type CityResponse struct {
	Count  int           `json:"count"`
	Cities []models.City `json:"cities"`
}

func NewCityHandler(service *services.CityService) *CityHandler {
	return &CityHandler{Service: service}
}

func (h *CityHandler) GetCities(c *fiber.Ctx) error {
	var cities []models.City
	for cityName, cityID := range models.CityCodes {
		cities = append(cities, models.City{
			ID:   cityID,
			City: cityName,
		})
	}

	// Şehir sayısı
	cityCount := len(cities)

	// JSON formatında döndürelim
	response := CityResponse{
		Count:  cityCount,
		Cities: cities,
	}

	return c.JSON(response)
}
