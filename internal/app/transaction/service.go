package transaction

type TransactionService interface {
	CreateTransaction(input CreateTransactionInput) (CreateTransactionOutput, error)
	GetAllTransactions() ([]GetAllTransactions, error)
}

type transactionService struct{
	transactionRepository TransactionRepository
}

func NewTransactionService(transactionRepository TransactionRepository) TransactionService {
	return &transactionService{
		transactionRepository: transactionRepository,
	}
}

func (s *transactionService) CreateTransaction(input CreateTransactionInput) (CreateTransactionOutput, error) {
	transaction := NewTransaction(input.Amount, input.Description, input.TransactionType)
	err := s.transactionRepository.Save(transaction)
	if err != nil {
		return CreateTransactionOutput{}, err
	}
	output := CreateTransactionOutput{
		ID: transaction.ID,
		Amount: transaction.Amount,
		Description: transaction.Description,
		CreatedAt: transaction.CreatedAt,
	}
	return output, nil
}

func (s *transactionService) GetAllTransactions() ([]GetAllTransactions, error) {
	transactions, err := s.transactionRepository.GetAll()
	if err != nil {
		return nil, err
	}
	var output []GetAllTransactions
	for _, transaction := range transactions {
		output = append(output, GetAllTransactions{
			ID: transaction.ID,
			Amount: transaction.Amount,
			Description: transaction.Description,
			CreatedAt: transaction.CreatedAt,
		})
	}
	return output, nil
}