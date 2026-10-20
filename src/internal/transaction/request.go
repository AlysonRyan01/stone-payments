package transaction

import (
	"errors"
	"strings"
	"time"
)

var (
	ErrInvalidPayload    = errors.New("invalid payload")
	ErrTimestampNotValid = errors.New("timestamp not valid")
	ErrTimestampOnFuture = errors.New("timestamp on the future")
)

type CreateTransactionRequest struct {
	CardNumber string `json:"cardNumber"`
	Amount     int64  `json:"amount"`
	Currency   string `json:"currency"`
	Merchant   string `json:"merchant"`
	Timestamp  string `json:"timestamp"`
}

func (r *CreateTransactionRequest) Validate() error {
	if strings.TrimSpace(r.CardNumber) == "" ||
		strings.TrimSpace(r.Currency) == "" ||
		strings.TrimSpace(r.Merchant) == "" ||
		r.Amount <= 0 {
		return ErrInvalidPayload
	}

	ts, err := time.Parse(time.RFC3339, r.Timestamp)
	if err != nil {
		return ErrTimestampNotValid
	}

	if ts.After(time.Now()) {
		return ErrTimestampOnFuture
	}

	return nil
}
