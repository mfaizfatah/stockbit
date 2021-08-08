package repository

import (
	"context"
	"database/sql"
	"project/app/model"
)

type repo struct {
	DB *sql.DB
}

type Repository interface {
	SaveMovie(ctx context.Context, searchWord string, s model.Search) (sql.Result, error)
}

func New(conn *sql.DB) Repository {
	return &repo{conn}
}
