package repository

import (
	"context"
	"database/sql"
	"project/app/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_repo_SaveMovie(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual)) // mock sql.DB
	assert.NoError(t, err)
	defer db.Close()

	r := New(db)

	type args struct {
		ctx        context.Context
		searchWord string
		s          model.Search
	}
	tests := []struct {
		name    string
		args    args
		want    sql.Result
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test insert",
			args: args{
				ctx:        context.TODO(),
				searchWord: "Batman",
				s: model.Search{
					Title:  "Batman: The Killing Joke",
					Year:   "2016",
					ImdbID: "tt4853102",
					Type:   "movie",
					Poster: "http://lorem.ipsum",
				},
			},
			want:    sqlmock.NewResult(1, 1),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := `INSERT INTO search(SearchWord, Title, Year, imdbID, Type, Poster) VALUES(?,?,?,?,?,?);`
			mock.ExpectExec(query).WithArgs(tt.args.searchWord, tt.args.s.Title, tt.args.s.Year, tt.args.s.ImdbID, tt.args.s.Type, tt.args.s.Poster).WillReturnResult(tt.want)

			_, err := r.SaveMovie(tt.args.ctx, tt.args.searchWord, tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("repo.SaveMovie() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
