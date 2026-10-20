package main

import (
	"database/sql"
	"fmt"
	"stone/internal/config"
	"stone/internal/transaction"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func BuildTransactionDependencies(db *sql.DB) *transaction.Router {
	repo := transaction.NewRepository(db)
	service := transaction.NewService(repo)
	handler := transaction.NewHandler(service)
	return transaction.NewRouter(handler)
}

func ConfigDB(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("pgx", cfg.ConnectionString())
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("ping db: %w", err)
	}
	return db, nil
}