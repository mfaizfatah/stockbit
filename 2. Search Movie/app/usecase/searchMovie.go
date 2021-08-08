package usecase

import (
	"context"
	"encoding/json"
	"net/http"
	"project/app/model"
	"sync"
)

func (r *uc) SearchMovie(c context.Context, searchWord, page string) (*model.OMDBResponse, int, error) {
	var result model.OMDBResponse
	var wg sync.WaitGroup

	omdbres, st, err := GetOMDB(searchWord, page)
	if err != nil {
		return nil, st, err
	}

	err = json.Unmarshal(omdbres, &result)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	wg.Add(len(result.Search))
	for i, value := range result.Search {
		go func(i int, value model.Search) {
			defer wg.Done()
			r.repo.SaveMovie(context.Background(), searchWord, value)
		}(i, value)
	}

	return &result, st, nil
}
