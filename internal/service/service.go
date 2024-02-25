package service

import (
	"log/slog"

	"github.com/olegtemek/file-handler/internal/model"
	"github.com/olegtemek/file-handler/internal/repository"
	"github.com/olegtemek/file-handler/internal/service/file"
)

type FileService interface {
	Create(filepath string, tag string) (*model.File, error)
	GetAll(params map[string]string) ([]*model.File, error)
	GetOne(id string) (*model.File, error)
	Delete(str string) (*model.File, error)
	GetAllTags() ([]*string, error)
}

type Service struct {
	FileService
}

func NewService(log *slog.Logger, repositories *repository.Repository) *Service {
	return &Service{
		FileService: file.NewService(log, &repositories.FileRepository),
	}
}
