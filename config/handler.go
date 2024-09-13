package config

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	Fiber *fiber.App
	config *Config	
}

func NewApp(config *Config) *App {
	app := fiber.New()
	return &App{
		Fiber: app,
		config: config,
	}
}

func (a *App) SetupRoutes() {
	a.Fiber.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hallo boskuhh semuanyaa.....")
	})
}

func (a *App) Start() error {
	return a.Fiber.Listen(fmt.Sprintf(":%s", a.config.AppPort))
}