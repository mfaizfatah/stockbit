package pb

import (
	"context"
	"project/app/model"
	"sync"
)

func (r *Transport) SearchMovie(ctx context.Context, req *SearchMovieRequest) (*OMDBResponse, error) {
	var wg sync.WaitGroup
	res, _, err := r.Uc.SearchMovie(ctx, req.Searchword, req.Pagination)
	if err != nil {
		return nil, err
	}

	search := make([]*Search, len(res.Search))

	wg.Add(len(res.Search))
	for i, vals := range res.Search {
		go func(i int, vals model.Search) {
			defer wg.Done()
			search[i] = &Search{
				Title:  vals.Title,
				Year:   vals.Year,
				ImdbID: vals.ImdbID,
				Type:   vals.Type,
				Poster: vals.Poster,
			}
		}(i, vals)
	}
	wg.Wait()

	return &OMDBResponse{
		Search:       search,
		Response:     res.Response,
		Error:        res.Error,
		TotalResults: res.TotalResults,
	}, nil
}
