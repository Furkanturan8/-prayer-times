package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"namaz-vakitleri/database"
	"namaz-vakitleri/handlers"
	"namaz-vakitleri/pkg/config"
	"namaz-vakitleri/routes"
	"namaz-vakitleri/services"
	"time"
)

func main() {
	fmt.Println("\n--------------BİSMİLLAH--------------\n")

	// Yapılandırma dosyasını yükle
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	app := fiber.New()

	// Middleware to log requests
	app.Use(requestLogger)

	// CORS middleware'i ekleyin
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:8080", // Vue uygulamanızın çalıştığı adres
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// DB Init
	db, err := database.DBInstance(cfg)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println("DB Address: ", db)

	// API Base URL
	apiBaseURL := "https://api.aladhan.com/v1/calendarByCity/2024/7?country=turkey&city="

	prayerTimesService := services.NewPrayerTimeService(apiBaseURL)
	cityService := services.NewCityService()
	phraseService := services.NewPhraseService(db)
	contactService := services.NewContactService(db)

	prayerTimesHandler := handlers.NewPrayerTimeHandler(prayerTimesService)
	cityHandler := handlers.NewCityHandler(cityService)
	phraseHandler := handlers.NewPhraseHandler(phraseService)
	contactHandler := handlers.NewContactHandler(contactService, cfg)

	routes.PhraseRoutes(app, phraseHandler)
	routes.CityRoutes(app, cityHandler)
	routes.PrayerTimeRoutes(app, prayerTimesHandler)
	routes.ContactRoutes(app, contactHandler)

	app.Listen(":3000")
}

func requestLogger(c *fiber.Ctx) error {
	start := time.Now()

	// Proceed to the next middleware or handler
	err := c.Next()

	stop := time.Now()
	latency := stop.Sub(start)

	// Get the status code and method
	status := c.Response().StatusCode()
	method := c.Method()
	url := c.OriginalURL()

	// Get the client IP
	clientIP := c.IP()

	// Log format: time | status | latency | clientIP | method | url
	fmt.Printf("%s | %3d | %9v | %15s | %-7s | %s\n",
		start.Format("15:04:05"),
		status,
		latency,
		clientIP,
		method,
		url,
	)

	return err
}
