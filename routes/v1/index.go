package v1

import (
	"fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

func Run(r *fiber.App) {
	getRoutes(r)
	r.Listen(":3000")
}

func getRoutes(app *fiber.App) {
	site := app.Group("/")
	api := app.Group("/api/v1", middleware.VerifyToken)
	TodosRoutes(api)
	SiteRoutes(site)
	AuthRoutes(api)
}
