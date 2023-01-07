package main

import (
	// v1 "fiber/routes/v1"
	v1 "fiber/routes/v1"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
)

func main() {
	engine := html.New("./views", ".html")
	// engine := handlebars.New("./views", ".hbs")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "./public")
	app.Use(cors.New())
	app.Use(logger.New())
	v1.Run(app)
	// v2.Run(app)
}
