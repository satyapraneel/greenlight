package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port string
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {

	var cfg config
	flag.StringVar(&cfg.port, "port", "4000", "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Env development|staging|production")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

	srv := &http.Server{
		Addr:         ":" + cfg.port,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	logger.Printf("Server started at port %s", cfg.port)
	err := srv.ListenAndServe()
	if err != nil {
		logger.Fatal(err)
	}
}
