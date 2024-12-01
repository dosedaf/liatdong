package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"liatdong/internal/delivery/http"
	"liatdong/internal/delivery/http/route"
	"liatdong/internal/repository"
	"liatdong/internal/usecase"
)

type ConfigInstance struct {
	App      *fiber.App
	Pool     *pgxpool.Pool
	Log      *logrus.Logger
	Config   *viper.Viper
	Validate *validator.Validate
}

func Bootstrap(config *ConfigInstance) {
	// repo
	fileRepository := repository.NewFileRepository(config.Pool, config.Log)

	// usecase
	fileUsecase := usecase.NewFileUsecase(config.Log, fileRepository)

	// controller
	fileController := http.NewFileController(fileUsecase, config.Log)

	routeConfig := route.RouteConfig{
		App:            config.App,
		FileController: fileController,
	}

	routeConfig.SetupRoutes()
}
