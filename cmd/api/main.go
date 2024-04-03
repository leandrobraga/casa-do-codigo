package main

import (
	"cdc/internal/data"
	"context"
	"database/sql"
	"flag"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type config struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
}

type application struct {
	config config
	models data.Models
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.db.dsn, "db-dsn", "postgres://leandro:123@db/casacodigo?sslmode=disable", "PostgreSQL DSN")

	flag.Parse()

	db, err := openDB(cfg)
	if err != nil {
		log.Fatalln(err, nil)
	}
	defer db.Close()

	app := &application{
		config: cfg,
		models: data.NewModels(db),
	}

	err = app.serve()
	if err != nil {
		log.Fatalln(err)
	}
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("pgx", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
