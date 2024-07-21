package routes

import (
	"github.com/gofiber/fiber/v2"
	"namaz-vakitleri/handlers"
)

func CityRoutes(app *fiber.App, cityHandler *handlers.CityHandler) {
	app.Get("/cities", cityHandler.GetCities)
}

func PrayerTimeRoutes(app *fiber.App, prayerTimeHandler *handlers.PrayerTimeHandler) {
	app.Get("/prayer-times/:city", prayerTimeHandler.GetPrayerTimesByCities)
}
