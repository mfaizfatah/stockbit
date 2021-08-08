package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *tp) SearchMovie(c *gin.Context) {
	res, st, err := r.uc.SearchMovie(c, c.Query("searchword"), c.Query("pagination"))
	if err != nil || st >= http.StatusBadRequest {
		c.JSON(st, err)
	}
	c.JSON(st, res)
}
