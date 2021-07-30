package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/workfoxes/gobase"
	"github.com/workfoxes/gobase/pkg/log"
)

func main() {
	app := gobase.New(&gobase.ApplicationConfig{Name: "TestingApplication", Port: 8080})
	app.Server.Get("/", func(c *fiber.Ctx) error {
		log.L.Info("Route for new service")
		log.L.Info("Route for new service")
		return c.SendString("Hello, Community")
	})
	app.Start()
}
