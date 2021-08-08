package route

import (
	"fmt"
	"log"

	"project/app/config"
	"project/app/delivery/http"

	"github.com/gin-gonic/gin"
)

type route struct {
	hd http.Transport
}

type Router interface {
	Router(port string)
}

func NewRoute(c http.Transport) Router {
	return &route{hd: c}
}

func (r *route) Router(port string) {
	// set gin mode based on environment
	switch config.AppEnv.APPENV {
	case "production", "staging":
		gin.SetMode(gin.ReleaseMode)
	case "development":
		gin.SetMode(gin.DebugMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	// init new gin
	routes := gin.New()
	routes.Use(gin.Recovery())

	// endpoint
	routes.GET("/", r.hd.SearchMovie)

	// run server
	log.Println(fmt.Sprintf("0.0.0.0%v 🚀 HTTP server started... ", port))
	routes.Run(port)
}
