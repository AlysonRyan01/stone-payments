package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"stone/internal/config"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := ConfigDB(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rdb, err := ConfigRedis(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer rdb.Close()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	txRouter := BuildTransactionDependencies(db, rdb)
	r.Route("/transactions", txRouter.Register)

	http.ListenAndServe(":3000", r)
}
