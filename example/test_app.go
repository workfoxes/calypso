package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/workfoxes/calypso"
	"github.com/workfoxes/calypso/pkg/log"
)

func main() {
	app := calypso.New(&calypso.ApplicationConfig{Name: "TestingApplication", Port: 8081})
	app.Server.Get("/", func(c *fiber.Ctx) error {
		log.Info("Route for new service")
		log.Info("Route for new service")
		return c.SendString("Hello, Community")
	})
	app.Start()
}
