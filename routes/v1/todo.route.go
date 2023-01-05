package v1

import (
	"fiber/controllers"

	"github.com/gofiber/fiber/v2"
)

func TodosRoutes(rg fiber.Router) {
	todos := rg.Group("/todos")
	todos.Get("/", controllers.GetTodos)
}
