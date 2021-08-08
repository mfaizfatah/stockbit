package http

import (
	"project/app/usecase"

	"github.com/gin-gonic/gin"
)

type tp struct {
	uc usecase.Usecase
}

type Transport interface {
	SearchMovie(c *gin.Context)
}

func NewTransport(u usecase.Usecase) Transport {
	return &tp{uc: u}
}
