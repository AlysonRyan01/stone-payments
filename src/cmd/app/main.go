package main

import (
	"log"
	"net/http"
	"stone/internal/config"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := ConfigDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	txRouter := BuildTransactionDependencies(db)
	r.Route("/transactions", txRouter.Register)

	http.ListenAndServe(":3000", r)
}
