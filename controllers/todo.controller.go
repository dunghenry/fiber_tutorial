package controllers

import "github.com/gofiber/fiber/v2"

func GetTodos(c *fiber.Ctx) error {
	return c.JSON((&fiber.Map{
		"status":  "success",
		"message": "Hello",
	}))
}
