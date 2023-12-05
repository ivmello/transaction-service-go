package transaction

type TransactionRepository interface {
	Save(transaction *Transaction) error
	GetAll() ([]Transaction,error)
}