package handlers

import (
	"github.com/gofiber/fiber/v2"
	"namaz-vakitleri/services"
	"strconv"
)

type PhraseHandler struct {
	Services *services.PhraseService
}

func NewPhraseHandler(service *services.PhraseService) *PhraseHandler {
	return &PhraseHandler{Services: service}
}

func (h *PhraseHandler) GetPhrases(c *fiber.Ctx) error {
	phrases, total, err := h.Services.GetPhrases()
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Not found")
	}
	return c.JSON(fiber.Map{
		"phrases":     phrases,
		"total_count": total,
	})
}

func (h *PhraseHandler) GetPhraseByID(c *fiber.Ctx) error {
	param := c.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return err
	}

	phrases, err := h.Services.GetPhraseByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Not found")
	}
	return c.JSON(fiber.Map{
		"phrases": phrases,
	})
}
