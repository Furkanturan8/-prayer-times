package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"namaz-vakitleri/database"
	"namaz-vakitleri/handlers"
	"namaz-vakitleri/routes"
	"namaz-vakitleri/services"
)

func main() {
	fmt.Println("\n--------------BİSMİLLAH--------------\n")

	app := fiber.New()

	// DB Init
	db, err := database.DBInstance()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("DB Address: ", db)

	// API Base URL
	apiBaseURL := "https://api.aladhan.com/v1/calendarByCity/2024/7?country=turkey&city="

	prayerTimesService := services.NewPrayerTimeService(apiBaseURL)
	cityService := services.NewCityService()

	prayerTimesHandler := handlers.NewPrayerTimeHandler(prayerTimesService)
	cityHandler := handlers.NewCityHandler(cityService)

	routes.CityRoutes(app, cityHandler)
	routes.PrayerTimeRoutes(app, prayerTimesHandler)

	app.Listen(":3000")
}
