package http

import (
	"net/http"
	"net/http/httptest"
	"project/app/repository"
	"project/app/usecase"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_tp_SearchMovie(t *testing.T) {
	db, _, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual)) // mock sql.DB
	assert.NoError(t, err)
	defer db.Close()

	repo := repository.New(db)
	uc := usecase.NewUC(repo)
	r := NewTransport(uc)

	tests := []struct {
		name         string
		url          string
		wantErr      bool
		wantRespCode int
	}{
		// TODO: Add test cases.
		{
			name:         "success",
			url:          "http://localhost:8080/?searchword=batman&pagination=2",
			wantErr:      false,
			wantRespCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		router := gin.Default()
		router.GET("/", r.SearchMovie)

		req, err := http.NewRequest(http.MethodGet, tt.url, nil)
		assert.NoError(t, err)

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		assert.Equal(t, resp.Code, tt.wantRespCode)
	}
}
