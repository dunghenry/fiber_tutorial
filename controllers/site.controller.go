package controllers

import (
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetHomePage(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Home Page",
	}, "layouts/main")
}

func UploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON((&fiber.Map{
			"status":  "success",
			"message": err.Error(),
		}))
	}
	filename := filepath.Base(file.Filename)
	data := time.Now().Unix()
	name := strconv.Itoa(int(data)) + "." + strings.Split(filename, ".")[1]
	if err := c.SaveFile(file, "public/images/"+name); err != nil {
		return c.JSON((&fiber.Map{
			"status":  "success",
			"message": err.Error(),
		}))
	}
	return c.SendString("File uploaded successfully.")
}
