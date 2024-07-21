package handlers

import (
	"github.com/gofiber/fiber/v2"
	"namaz-vakitleri/models"
	"namaz-vakitleri/services"
)

type PrayerTimeHandler struct {
	Service *services.PrayerTimeService
}

func NewPrayerTimeHandler(service *services.PrayerTimeService) *PrayerTimeHandler {
	return &PrayerTimeHandler{Service: service}
}

func (h *PrayerTimeHandler) GetPrayerTimesByCities(c *fiber.Ctx) error {
	cityName := c.Params("city")

	cityID, exists := models.CityCodes[cityName]
	if !exists {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid city name")
	}

	city := models.City{
		ID:   cityID,
		City: cityName,
	}

	prayerTimes, err := h.Service.GetPrayerTimesByCity(city)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Not found")
	}
	return c.JSON(prayerTimes)
}
