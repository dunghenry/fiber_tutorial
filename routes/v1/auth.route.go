package v1

import (
	"fiber/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(rg fiber.Router) {
	auth := rg.Group("/auth")
	auth.Post("/register", controllers.Register)
	auth.Post("/login", controllers.Login)
}
