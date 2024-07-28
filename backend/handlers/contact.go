package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"namaz-vakitleri/helpers"
	"namaz-vakitleri/models"
	"namaz-vakitleri/pkg/config"
	"namaz-vakitleri/services"
)

type ContactHandler struct {
	Services *services.ContactService
	Config   *config.Config
}

func NewContactHandler(service *services.ContactService, cfg *config.Config) *ContactHandler {
	return &ContactHandler{
		Services: service,
		Config:   cfg,
	}
}

func (h *ContactHandler) Create(c *fiber.Ctx) error {
	var contact models.Contact
	if err := c.BodyParser(&contact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz JSON formatı"})
	}

	err := h.Services.Create(contact)
	if err != nil {
		log.Printf("Mesaj gönderilirken hata oluştu: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Mesaj oluşturulurken hata oluştu"})
	}

	err = helpers.SendMailToAdmin(h.Config, contact.Email, contact.Name, contact.Surname, contact.Email, contact.Message)
	if err != nil {
		fmt.Println("E-posta gönderilemedi:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Email gönderilirken hata oluştu"})
	}

	return c.SendString("Mesaj gönderildi!")
}
