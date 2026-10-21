package transaction

import (
	"context"
	"database/sql"
	"fmt"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(ctx context.Context, t *Transaction) (string, error) {
	var ID string
	err := r.db.QueryRowContext(ctx, `
		INSERT INTO transaction (card_number, amount, currency, merchant, timestamp)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING ID`, t.CardNumber, t.Amount, t.Currency, t.Merchant, t.Timestamp,).Scan(&ID)

	if err != nil {
		return "", err
	}

	return ID, nil
}

func (r *Repository) CreateSuspiciousTransaction(ctx context.Context, t *Transaction, reason string) (string, error) {
	var transactionID string
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return "", err
	}
	defer tx.Rollback()

	err = tx.QueryRowContext(ctx, `
		INSERT INTO transaction (card_number, amount, currency, merchant, timestamp)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING ID`, t.CardNumber, t.Amount, t.Currency, t.Merchant, t.Timestamp,).Scan(&transactionID)
	if err != nil {
		return "", fmt.Errorf("insert transaction: %w", err)
	}
	
	_, err = tx.ExecContext(ctx, `
		INSERT INTO suspicious_transaction (transaction_id, reason)
		VALUES ($1, $2)`, transactionID, reason)
	if err != nil {
		return "", fmt.Errorf("insert suspicious: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return "", fmt.Errorf("commit: %w", err)
	}

	return transactionID, nil
}