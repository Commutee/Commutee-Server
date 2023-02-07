package routes

import (
	"commutee/handler"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Matrix(c *fiber.Ctx) error {
	l := c.Query("start")
	k := c.Query("end")
	mode := c.Query("mode")
	if k == "" || l == "" {
		return c.SendString("404")
	}
	if mode == "" {
		mode = "bus"
	}

	start, err := handler.CoodinateCalc(l)
	if err != nil {
		log.Fatal(err)
	}
	end, err := handler.CoodinateCalc(k)
	if err != nil {
		log.Fatal(err)
	}

	instructions, err := handler.WaypointCalc(fmt.Sprintf("%s|%s", start, end), mode)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server error")
	}
	return c.JSON(instructions)
}
