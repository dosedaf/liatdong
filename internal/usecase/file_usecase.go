package usecase

import (
	"github.com/sirupsen/logrus"

	"liatdong/internal/repository"
)

type FileUsecase struct {
	Log  *logrus.Logger
	Repo *repository.FileRepository
}

func NewFileUsecase(log *logrus.Logger, repo *repository.FileRepository) *FileUsecase {
	return &FileUsecase{
		Log:  log,
		Repo: repo,
	}
}
