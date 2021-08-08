package usecase

import (
	"context"
	"net/http"
	"project/app/repository"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_uc_SearchMovie(t *testing.T) {
	db, _, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual)) // mock sql.DB
	assert.NoError(t, err)
	defer db.Close()

	repo := repository.New(db)
	r := NewUC(repo)

	type args struct {
		c          context.Context
		searchWord string
		page       string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				c:          context.TODO(),
				searchWord: "Batman",
				page:       "2",
			},
			want:    "True",
			want1:   http.StatusOK,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := r.SearchMovie(tt.args.c, tt.args.searchWord, tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("uc.SearchMovie() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Response, tt.want) {
				t.Errorf("uc.SearchMovie() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("uc.SearchMovie() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
