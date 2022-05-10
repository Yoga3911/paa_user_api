package main

import (
	"log"
	"os"
	"user_api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routes.Route(app)
	log.Fatal(app.Listen(":"+ os.Getenv("PORT")))
}
