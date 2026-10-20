package transaction

import "time"

type Transaction struct {
	ID string
	CardNumber string
	Amount int64
	Currency string
	Merchant string
	Timestamp time.Time
}

