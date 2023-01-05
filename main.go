package main

import (
	// v1 "fiber/routes/v1"
	v2 "fiber/routes/v2"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	//v1.Run(app)
	v2.Run(app)
}
