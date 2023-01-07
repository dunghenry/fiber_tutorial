package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetTodos(c *fiber.Ctx) error {
	fmt.Println(c.GetReqHeaders()["token"])
	return c.JSON((&fiber.Map{
		"status":  "success",
		"message": "Hello",
	}))
}
