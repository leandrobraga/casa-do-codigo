package main

import (
	"flag"
	"log"
)

type config struct {
	port int
	env  string
}

type application struct {
	config config
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	app := &application{
		config: cfg,
	}

	err := app.serve()
	if err != nil {
		log.Fatalln(err)
	}
}
