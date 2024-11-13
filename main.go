package main

import (
	"fmt"
	"log"
	"os"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/quad-x/awb-sls/routes"
)

func main() {
	api := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	api.Use(cors.New())

	api.Get("/", routes.GetHealth)
	api.Post("/events", routes.ReceiveEvents)
	api.Post("/consolidated", routes.GenerateConsolidatedAWB)

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	log.Fatal(api.Listen(port))
}
