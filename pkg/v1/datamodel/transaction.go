package datamodel

type Transaction struct {
	Amount          float64 `json:"amount"`
	Message         string  `json:"message"`
	TransactionType string  `json:"transaction_type"`
	UserUUID        string  `json:"user_uuid"`
}
