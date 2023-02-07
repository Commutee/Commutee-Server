package routes

import (
	"commutee/handler"

	"github.com/gofiber/fiber/v2"
)

func Recomendations(c *fiber.Ctx) error {
	content, err := handler.LoadRecomendations()
	if err != nil {
		return c.SendString("Error")
	}
	return c.JSON(content)
}
