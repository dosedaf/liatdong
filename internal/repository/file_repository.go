package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type FileRepository struct {
	Log  *logrus.Logger
	Pool *pgxpool.Pool
}

func NewFileRepository(pool *pgxpool.Pool, log *logrus.Logger) *FileRepository {
	return &FileRepository{
		Pool: pool,
		Log:  log,
	}
}

func (f *FileRepository) Get() {
}
