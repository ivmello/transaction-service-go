package database

import (
	"database/sql"
	"transaction-service-go/internal/app/transaction"
)

type postgresAdapter struct {
	db *sql.DB
}

func NewPostgresAdapter(db *sql.DB) transaction.TransactionRepository {
	return &postgresAdapter{db: db}
}

func (m *postgresAdapter) Save(transaction *transaction.Transaction) error {
	stmt, err := m.db.Prepare("INSERT INTO transactions(id, amount, description, transaction_type, created_at) VALUES($1, $2, $3, $4, $5)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(transaction.ID, transaction.Amount, transaction.Description, transaction.TransactionType, transaction.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (m *postgresAdapter) GetAll() ([]transaction.Transaction, error) {
	var transactions []transaction.Transaction
	rows, err := m.db.Query("SELECT id, amount, description, transaction_type, created_at FROM transactions ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var transaction transaction.Transaction
		err := rows.Scan(&transaction.ID, &transaction.Amount, &transaction.Description, &transaction.TransactionType, &transaction.CreatedAt)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}