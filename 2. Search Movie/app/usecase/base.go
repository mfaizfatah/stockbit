package usecase

import (
	"context"
	"project/app/model"
	"project/app/repository"
)

type uc struct {
	repo repository.Repository
}

type Usecase interface {
	SearchMovie(c context.Context, searchWord, page string) (*model.OMDBResponse, int, error)
}

func NewUC(repo repository.Repository) Usecase {
	return &uc{repo}
}
