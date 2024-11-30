package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ConfigInstance struct {
	App      *fiber.App
	DB       *pgx.Conn
	Log      *logrus.Logger
	Config   *viper.Viper
	Validate *validator.Validate
}

func Bootstrap(config *ConfigInstance) {
}
