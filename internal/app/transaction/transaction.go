package transaction

import (
	"time"

	"github.com/google/uuid"
)

type TransactionType string
const (
	Credit TransactionType = "credit"
	Debit TransactionType = "debit"
)

type Transaction struct {
	ID string 
	Amount int
	Description string
	TransactionType TransactionType
	CreatedAt time.Time
}

func NewTransaction(amount int, description string, transactionType TransactionType) *Transaction {
	return &Transaction{
		ID: uuid.New().String(),
		Amount: amount,
		Description: description,
		TransactionType: transactionType,
		CreatedAt: time.Now(),
	}
}