package v1

import (
	"fiber/controllers"

	"github.com/gofiber/fiber/v2"
)

func SiteRoutes(rg fiber.Router) {
	site := rg.Group("/")
	site.Get("/", controllers.GetHomePage)
	site.Post("/upload", controllers.UploadFile)
}
