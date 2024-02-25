package file

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/olegtemek/file-handler/internal/dto"
	"github.com/olegtemek/file-handler/internal/response"
	"github.com/olegtemek/file-handler/internal/service"
	"github.com/olegtemek/file-handler/pkg/file"
)

type Handler struct {
	log     *slog.Logger
	service *service.FileService
}

func NewHandler(log *slog.Logger, service *service.FileService) *Handler {
	return &Handler{
		log:     log,
		service: service,
	}
}

// @Summary create
// @Tags file
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File"
// @Param tag formData string true "tag name"
// @Success 200 {object} response.FileCreate
// @Router /file [post]
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	h.log = h.log.With("Source", "FileHandler:Create")

	errParseSize := r.ParseMultipartForm(32 << 20) // 32 MiB max size
	if errParseSize != nil {
		response.NewError(&w, r, fmt.Errorf("max file size 32 MiB. %w", errParseSize), 400)
		return
	}

	gettingFile, handler, err := r.FormFile("file")

	if err != nil {
		response.NewError(&w, r, fmt.Errorf("cannot get file. %w", err), 400)
		return
	}

	var req dto.FileCreateDto

	req.Tag = r.FormValue("tag")
	req.File = &gettingFile

	if err := validator.New().Struct(req); err != nil {
		validateErr := err.(validator.ValidationErrors)
		render.JSON(w, r, response.ValidationError(validateErr))
		return
	}

	filepath, err := file.SaveFile(&gettingFile, handler)

	if err != nil {
		response.NewError(&w, r, fmt.Errorf("cannot store file. %w", err), 400)
		return
	}

	file, err := (*h.service).Create(filepath, req.Tag)
	if err != nil {
		response.NewError(&w, r, fmt.Errorf("cannot save file to db. %w", err), 400)
		return
	}

	render.JSON(w, r, &response.FileCreate{
		Status: 200,
		File:   file,
	})
}

// @Summary getAll
// @Tags file
// @Accept json
// @Produce json
// @Param q query string false "Tag"
// @Param q query string false "Filename"
// @Success 200 {object} response.FileGetAll
// @Router /file [get]
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	params := make(map[string]string)

	if r.URL.Query().Get("filepath") != "" {
		params["filepath"] = r.URL.Query().Get("filepath")
	}
	if r.URL.Query().Get("tag") != "" {
		params["tag"] = r.URL.Query().Get("tag")
	}

	files, err := (*h.service).GetAll(params)
	if err != nil {
		response.NewError(&w, r, fmt.Errorf("cannot get files from db. %w", err), 400)
		return
	}

	render.JSON(w, r, &response.FileGetAll{
		Status: 200,
		Files:  files,
	})
}

// @Summary getOne
// @Tags file
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} response.FileGetAll
// @Router /file/{id} [get]
func (h *Handler) GetOne(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	file, err := (*h.service).GetOne(id)
	if err != nil {
		if err == sql.ErrNoRows {
			response.NewError(&w, r, fmt.Errorf("file not found. %w", err), 404)
			return
		}
		response.NewError(&w, r, fmt.Errorf("cannot get file from db %w", err), 400)
		return
	}

	render.JSON(w, r, &response.FileGetOne{
		Status: 200,
		File:   file,
	})
}

// @Summary delete
// @Tags file
// @Accept json
// @Produce json
// @Param id path string false "Tag"
// @Success 200 {object} response.FileDelete
// @Router /file/{id} [delete]
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	deletedFile, err := (*h.service).Delete(id)
	if err != nil {
		if err == sql.ErrNoRows {
			response.NewError(&w, r, fmt.Errorf("file not found. %w", err), 404)
			return
		}
		response.NewError(&w, r, fmt.Errorf("cannot delete file from db %w", err), 400)
		return
	}

	err = file.DeleteFile(deletedFile.FilePath)
	if err != nil {
		response.NewError(&w, r, fmt.Errorf("cannot delete file from storage %w", err), 400)
		return
	}

	render.JSON(w, r, &response.FileDelete{
		Status: 200,
		File:   deletedFile,
	})
}
