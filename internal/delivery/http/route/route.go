package route

import (
	"github.com/gofiber/fiber/v2"

	"liatdong/internal/delivery/http"
)

type RouteConfig struct {
	App            *fiber.App
	FileController *http.FileController
}

func (c *RouteConfig) SetupRoutes() {
	c.App.Get("/upload", c.FileController.Upload)
}
