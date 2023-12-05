package transaction

import "time"

type IError struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

type CreateTransactionInput struct {
	Amount int `json:"amount" validate:"required,numeric"`
	Description string `json:"description" validate:"required,max=100"`
	TransactionType TransactionType `json:"transaction_type" validate:"required,oneof=credit debit"`
}

type CreateTransactionOutput struct {
	ID string `json:"id"`
	Amount int `json:"amount"`
	Description string `json:"description"`
	TransactionType TransactionType `json:"transaction_type"`
	CreatedAt time.Time `json:"created_at"`
}

type GetAllTransactions struct {
	ID string `json:"id"`
	Amount int `json:"amount"`
	Description string `json:"description"`
	TransactionType TransactionType `json:"transaction_type"`
	CreatedAt time.Time `json:"created_at"`
}