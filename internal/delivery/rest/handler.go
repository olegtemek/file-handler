package rest

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	_ "github.com/olegtemek/file-handler/docs"
	"github.com/olegtemek/file-handler/internal/config"
	"github.com/olegtemek/file-handler/internal/delivery/rest/file"
	"github.com/olegtemek/file-handler/internal/middleware/logger"
	"github.com/olegtemek/file-handler/internal/service"
	httpSwagger "github.com/swaggo/http-swagger"
)

type FileHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetOne(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetAllTags(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	Log     *slog.Logger
	Cfg     *config.Config
	Handler *chi.Mux
	FileHandler
}

func NewHandler(log *slog.Logger, cfg *config.Config, services *service.Service) *Handler {
	return &Handler{
		Log:         log,
		Cfg:         cfg,
		Handler:     chi.NewRouter(),
		FileHandler: file.NewHandler(log, &services.FileService),
	}
}

func (h *Handler) Init() *http.Server {

	h.Handler.Use(middleware.RequestID)
	h.Handler.Use(middleware.URLFormat)
	h.Handler.Use(middleware.Recoverer)
	h.Handler.Use(logger.New(h.Log))
	h.Handler.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"POST", "GET", "DELETE"},
	}))

	h.Handler.Mount("/swagger", httpSwagger.WrapHandler)

	fs := http.FileServer(http.Dir("uploads"))

	h.Handler.Handle("/uploads/*", http.StripPrefix("/uploads/", fs))

	h.InitAllRoutes()

	handler := h.Handler

	srv := &http.Server{
		Addr:        h.Cfg.Address,
		ReadTimeout: h.Cfg.Timeout,
		Handler:     handler,
	}

	return srv
}

func (h *Handler) InitAllRoutes() {

	h.Handler.Route("/v1", func(v1 chi.Router) {
		v1.Route("/file", func(r chi.Router) {
			r.Post("/", h.FileHandler.Create)
			r.Get("/", h.FileHandler.GetAll)
			r.Get("/tags", h.FileHandler.GetAllTags)
			r.Get("/{id}", h.FileHandler.GetOne)
			r.Delete("/{id}", h.FileHandler.Delete)
		})
	})
}
