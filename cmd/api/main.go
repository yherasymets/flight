package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/yherasymets/flight/internal/repo"
)

const version = "1.0.0"

type db struct {
	dsn string
}

type config struct {
	port int
	env  string
	db
}

type app struct {
	config config
	logger *log.Logger
	repo   repo.Service
}

func main() {
	var config config

	flag.IntVar(&config.port, "port", 8000, "Server port")
	flag.StringVar(&config.env, "env", "development", "Environment (development/staging/production)")

	flag.StringVar(&config.db.dsn, "database-dsn", os.Getenv("FLIGHT_DB_DSN"), "Database DSN")
	flag.Parse()

	logger := log.New(os.Stdout, "# ", log.Ldate|log.Ltime)

	db, err := connectionDB(config)
	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()

	logger.Printf("connected to database PostgreSQL")

	app := &app{
		config: config,
		logger: logger,
		repo:   repo.NewService(db),
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", config.env, srv.Addr)
	err = srv.ListenAndServe()
	logger.Fatal(err)
}

func connectionDB(cfg config) (*sql.DB, error) {

	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}
	// Create a context with a 5-second timeout deadline.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
