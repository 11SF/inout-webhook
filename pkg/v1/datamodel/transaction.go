package datamodel

type Transaction struct {
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	UserUUID        string  `json:"user_uuid"`
}
