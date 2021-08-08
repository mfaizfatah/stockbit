package main

import (
	"project/app/adapter"
	"project/app/config"
	"project/app/delivery/http"
	"project/app/repository"
	"project/app/route"
	"project/app/usecase"

	"github.com/gomodul/fn"
)

func init() {
	adapter.LoadMySQL()
}

func main() {
	db := adapter.DBSQL()
	defer fn.Check(db.Close)

	repository := repository.New(db)
	usecase := usecase.NewUC(repository)
	delivery := http.NewTransport(usecase)

	router := route.NewRoute(delivery)

	go router.Router(config.Server.PORTHTTP)
	route.PbClient(usecase)
}
