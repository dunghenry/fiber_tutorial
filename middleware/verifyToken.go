package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func VerifyToken(c *fiber.Ctx) error {
	x := "Hi"
	c.Request().Header.Add("token", x)
	fmt.Printf(x)
	return c.Next()
}
