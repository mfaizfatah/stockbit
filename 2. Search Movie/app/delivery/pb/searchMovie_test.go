package pb

import (
	"context"
	"project/app/repository"
	"project/app/usecase"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestTransport_SearchMovie(t *testing.T) {
	db, _, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual)) // mock sql.DB
	assert.NoError(t, err)
	defer db.Close()

	repo := repository.New(db)
	uc := usecase.NewUC(repo)
	r := Transport{
		Uc: uc,
	}

	type args struct {
		ctx context.Context
		req *SearchMovieRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *OMDBResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				ctx: context.TODO(),
				req: &SearchMovieRequest{
					Searchword: "Batman",
					Pagination: "2",
				},
			},
			want: &OMDBResponse{
				Response: "True",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.SearchMovie(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Transport.SearchMovie() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Response, tt.want.Response) {
				t.Errorf("Transport.SearchMovie() = %v, want %v", got.Response, tt.want.Response)
			}
		})
	}
}
