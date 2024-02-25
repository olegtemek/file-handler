package file

import (
	"log/slog"

	"github.com/olegtemek/file-handler/internal/model"
	"github.com/olegtemek/file-handler/internal/repository"
)

type Service struct {
	log        *slog.Logger
	repository *repository.FileRepository
}

func NewService(log *slog.Logger, repository *repository.FileRepository) *Service {
	return &Service{
		log:        log,
		repository: repository,
	}
}

func (s *Service) Create(filepath string, tag string) (*model.File, error) {
	s.log = s.log.With(slog.String("Source", "FileService:Create"))

	file, err := (*s.repository).Create(filepath, tag)
	if err != nil {
		s.log.Error("Cannot save file", err)
		return file, err
	}

	return file, nil
}

func (s *Service) GetAll(params map[string]string) ([]*model.File, error) {
	s.log = s.log.With(slog.String("Source", "FileService:GetAll"))

	files, err := (*s.repository).GetAll(params)
	if err != nil {
		s.log.Error("Cannot get files", err)
		return files, err
	}

	return files, nil
}
func (s *Service) GetAllTags() ([]*string, error) {
	s.log = s.log.With(slog.String("Source", "FileService:GetAllTags"))

	tags, err := (*s.repository).GetAllTags()
	if err != nil {
		s.log.Error("Cannot get tags", err)
		return tags, err
	}

	return tags, nil
}

func (s *Service) Delete(id string) (*model.File, error) {
	s.log = s.log.With(slog.String("Source", "FileService:Delete"))

	file, err := (*s.repository).Delete(id)
	if err != nil {
		s.log.Error("Cannot delete file", err)
		return file, err
	}

	return file, nil
}

func (s *Service) GetOne(id string) (*model.File, error) {
	s.log = s.log.With(slog.String("Source", "FileService:GetOne"))

	file, err := (*s.repository).Delete(id)
	if err != nil {
		s.log.Error("Cannot get file", err)
		return file, err
	}

	return file, nil
}
