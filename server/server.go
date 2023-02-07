package server

import (
	"commutee/routes"
	"commutee/types"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Initserver(config types.Config) {
	app := fiber.New()
	app.Get("/", routes.HelloWorld)
	app.Get("/api/recommendation", routes.Recomendations)
	app.Get("/api/matrix", routes.Matrix)
	log.Fatal(app.Listen(":" + strconv.Itoa(config.Server.Port)))
}
