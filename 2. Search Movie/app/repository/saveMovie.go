package repository

import (
	"context"
	"database/sql"
	"project/app/model"
)

const TodoCreateQuery = `INSERT INTO search(SearchWord, Title, Year, imdbID, Type, Poster) VALUES(?,?,?,?,?,?);`

func (r *repo) SaveMovie(ctx context.Context, searchWord string, s model.Search) (sql.Result, error) {
	return r.DB.ExecContext(ctx, TodoCreateQuery, searchWord, s.Title, s.Year, s.ImdbID, s.Type, s.Poster)
}
