package transaction

import (
	"context"
	"database/sql"
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