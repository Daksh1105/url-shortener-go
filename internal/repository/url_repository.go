package repository

import (
	"context"
	"database/sql"
)

type URLRepository interface {
	Save(ctx context.Context, originalURL string) (uint64, error)
	GetByID(ctx context.Context, id uint64) (string, error)
}

type postgresURLRepository struct {
	db *sql.DB
}

func NewPostgresURLRepository(db *sql.DB) URLRepository {
	return &postgresURLRepository{db: db}
}

func (r *postgresURLRepository) Save(ctx context.Context, originalURL string) (uint64, error) {
	query := `INSERT INTO urls (original_url) VALUES ($1) RETURNING id`
	var id uint64
	err := r.db.QueryRowContext(ctx, query, originalURL).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *postgresURLRepository) GetByID(ctx context.Context, id uint64) (string, error) {
	query := `SELECT original_url FROM urls WHERE id = $1`
	var originalURL string
	err := r.db.QueryRowContext(ctx, query, id).Scan(&originalURL)
	if err != nil {
		return "", err
	}
	return originalURL, nil
}
