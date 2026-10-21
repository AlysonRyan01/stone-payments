package transaction

type TransactionResponse struct {
	Status      string
	AuthorizeID string
	Warning     string
	Error       string
}
