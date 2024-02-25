package file

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/olegtemek/file-handler/internal/model"
)

type Repository struct {
	log *slog.Logger
	db  *pgxpool.Pool
}

func NewRepository(log *slog.Logger, db *pgxpool.Pool) *Repository {
	return &Repository{
		log: log,
		db:  db,
	}
}

func (fr *Repository) Create(filepath string, tag string) (*model.File, error) {
	fr.log = fr.log.With(slog.String("Source", "FileRepository:Create"))

	var file model.File

	q := `INSERT INTO File(filepath,tag) VALUES($1, $2) RETURNING *`
	fr.log.Info("", slog.String("DB QUERY", q))
	if err := fr.db.QueryRow(context.Background(), q, filepath, tag).Scan(&file.Id, &file.FilePath, &file.Tag, &file.Timestamp); err != nil {
		return &file, err
	}

	return &file, nil
}

func (fr *Repository) GetAll(params map[string]string) ([]*model.File, error) {
	fr.log = fr.log.With(slog.String("Source", "FileRepository:GetAll"))

	// var files []*model.File

	files := []*model.File{}

	whereClause := ""

	for key, value := range params {
		if key == "tag" {
			whereClause += fmt.Sprintf(" AND %s = '%s'", key, value)
		}
		if key == "filepath" {
			whereClause += fmt.Sprintf(" AND %s = '%s'", key, value)
		}
	}

	if whereClause != "" {
		whereClause = whereClause[4:]
		whereClause = fmt.Sprintf(" WHERE %s", whereClause)
	}

	q := fmt.Sprintf("SELECT * FROM File%s", whereClause)

	fr.log.Info("", slog.String("DB QUERY", q))

	rows, err := fr.db.Query(context.Background(), q)
	if err != nil {
		return files, err
	}
	defer rows.Close()

	for rows.Next() {
		var file model.File
		err := rows.Scan(&file.Id, &file.FilePath, &file.Tag, &file.Timestamp)
		if err != nil {
			return nil, err
		}
		files = append(files, &file)
	}

	return files, nil
}

func (fr *Repository) Delete(id string) (*model.File, error) {
	fr.log = fr.log.With(slog.String("Source", "FileRepository:Delete"))

	var file model.File

	q := `DELETE FROM File WHERE id = $1 RETURNING id, filepath, tag, timestamp`
	fr.log.Info("", slog.String("DB QUERY", q))
	if err := fr.db.QueryRow(context.Background(), q, id).Scan(&file.Id, &file.FilePath, &file.Tag, &file.Timestamp); err != nil {
		return &file, err
	}

	return &file, nil
}
func (fr *Repository) GetOne(id string) (*model.File, error) {
	fr.log = fr.log.With(slog.String("Source", "FileRepository:GetOne"))

	var file model.File

	q := `SELECT * FROM File WHERE id = $1`
	fr.log.Info("", slog.String("DB QUERY", q))
	if err := fr.db.QueryRow(context.Background(), q, id).Scan(&file.Id, &file.FilePath, &file.Tag, &file.Timestamp); err != nil {
		return &file, err
	}

	return &file, nil
}
