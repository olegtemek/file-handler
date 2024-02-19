package db

import (
	"context"
	"log/slog"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/olegtemek/file-handler/internal/config"
)

func NewPostgresConnect(log *slog.Logger, cfg *config.Config) (*pgxpool.Pool, error) {

	conn, err := pgxpool.New(context.Background(), cfg.DatabaseUrl)

	if err != nil {
		log.Error("cannot connect to db", err)
		return conn, err
	}

	return conn, nil
}
