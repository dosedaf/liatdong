package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func NewFiberApp(viper *viper.Viper) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:       viper.GetString("app.name"),
		CaseSensitive: true,
		StrictRouting: true,
	})

	return app
}
