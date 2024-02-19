package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/olegtemek/file-handler/internal/config"
	"github.com/olegtemek/file-handler/internal/delivery/rest"
	"github.com/olegtemek/file-handler/internal/logger"
	"github.com/olegtemek/file-handler/internal/repository"
	"github.com/olegtemek/file-handler/internal/service"
	"github.com/olegtemek/file-handler/pkg/db"
)

// @title File-Handler Backend
// @version 1.0

// @BasePath /v1
func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic("Cannot config file")
	}

	log := logger.NewLogger(cfg)

	db, err := db.NewPostgresConnect(log, cfg)
	defer db.Close()
	if err != nil {
		panic("Cannot connect to database")
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	repositories := repository.NewRepository(log, db)
	services := service.NewService(log, repositories)
	restHandler := rest.NewHandler(log, cfg, services)

	server := restHandler.Init()

	go func() {
		server.ListenAndServe()
	}()

	log.Info("server started")

	<-done

	log.Info("stopping server")

	time.Sleep(time.Second * 5)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Error("failed to stop server", err)

		return
	}

	log.Info("server stopped")
}
