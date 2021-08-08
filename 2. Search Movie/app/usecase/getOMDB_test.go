package usecase

import (
	"encoding/json"
	"net/http"
	"project/app/model"
	"reflect"
	"testing"

	"github.com/gomodul/envy"
	"github.com/stretchr/testify/assert"
)

func TestGetOMDB(t *testing.T) {
	envy.Set("APP_ENV", "test")
	type args struct {
		searchword string
		pagination string
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
				searchword: "Batman",
				pagination: "2",
			},
			want:    "True",
			want1:   http.StatusOK,
			wantErr: false,
		},
		{
			name: "blank searchword",
			args: args{
				searchword: "",
				pagination: "2",
			},
			want:    "False",
			want1:   http.StatusOK,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := GetOMDB(tt.args.searchword, tt.args.pagination)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOMDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var res model.OMDBResponse
			err = json.Unmarshal(got, &res)
			assert.NoError(t, err)

			if !reflect.DeepEqual(res.Response, tt.want) {
				t.Errorf("GetOMDB() got = %v, want %v", res.Response, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetOMDB() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
