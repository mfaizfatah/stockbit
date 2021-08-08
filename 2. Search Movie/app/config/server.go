package config

import "github.com/gomodul/envy"

type server struct {
	PORTHTTP  string
	PORTPROTO string
}

var Server = server{
	PORTHTTP:  ":" + envy.Get("PORT_HTTP", "80"),
	PORTPROTO: ":" + envy.Get("PORT_PROTO", "81"),
}
