package main

import (
	"context"
	"database/sql"
	"fmt"
	"stone/internal/config"
	"stone/internal/transaction"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/redis/go-redis/v9"
)

func BuildTransactionDependencies(db *sql.DB, redis *redis.Client) *transaction.Router {
	repo := transaction.NewRepository(db)
	velocity := transaction.NewRedisVelocity(redis)
	service := transaction.NewService(repo, velocity)
	handler := transaction.NewHandler(service)
	return transaction.NewRouter(handler)
}

func ConfigDB(ctx context.Context, cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("pgx", cfg.ConnectionString())
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("ping db: %w", err)
	}

	return db, nil
}

func ConfigRedis(ctx context.Context, cfg *config.Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr(),
		Password: cfg.RedisPassword,
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		rdb.Close()
		return nil, fmt.Errorf("ping redis: %w", err)
	}

	return rdb, nil
}