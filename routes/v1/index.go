package v1

import "github.com/gofiber/fiber/v2"

func Run(r *fiber.App) {
	getRoutes(r)
	r.Listen(":3000")
}

func getRoutes(app *fiber.App) {
	api := app.Group("/api/v1")
	TodosRoutes(api)
}
