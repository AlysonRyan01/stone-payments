package transaction

import "time"

func CreateRequestToModel(r CreateTransactionRequest) (*Transaction, error) {
	ts, err := time.Parse(time.RFC3339, r.Timestamp)
	if err != nil {
		return nil, ErrTimestampNotValid
	}

	return &Transaction{
		CardNumber: r.CardNumber,
		Amount:     r.Amount,
		Currency:   r.Currency,
		Merchant:   r.Merchant,
		Timestamp:  ts,
	}, nil
}