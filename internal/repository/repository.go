package repository

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/olegtemek/file-handler/internal/model"
	"github.com/olegtemek/file-handler/internal/repository/file"
)

type FileRepository interface {
	Create(filepath string, tag string) (*model.File, error)
	GetAll(params map[string]string) ([]*model.File, error)
	GetOne(id string) (*model.File, error)
	Delete(id string) (*model.File, error)
}

type Repository struct {
	FileRepository
}

func NewRepository(log *slog.Logger, db *pgxpool.Pool) *Repository {
	return &Repository{
		FileRepository: file.NewRepository(log, db),
	}
}
