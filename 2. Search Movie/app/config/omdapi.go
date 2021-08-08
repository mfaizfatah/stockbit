package config

import "github.com/gomodul/envy"

type omdbapi struct {
	Url    string
	ApiKey string
}

var OmdbVar = omdbapi{
	Url:    envy.Get("OMDB_API_URL", "http://www.omdbapi.com"),
	ApiKey: envy.Get("OMDB_APIKEY"),
}
