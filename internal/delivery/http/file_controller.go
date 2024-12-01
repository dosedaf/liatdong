package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	"liatdong/internal/usecase"
)

type FileController struct {
	Log     *logrus.Logger
	Usecase *usecase.FileUsecase
}

func NewFileController(usecase *usecase.FileUsecase, log *logrus.Logger) *FileController {
	return &FileController{
		Log:     log,
		Usecase: usecase,
	}
}

func (f *FileController) Upload(ctx *fiber.Ctx) error {
	return ctx.JSON("mantap")
}
