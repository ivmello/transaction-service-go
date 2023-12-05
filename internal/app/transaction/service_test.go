package transaction

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockTransactionRepository struct{
	mock.Mock
}

func (m *mockTransactionRepository) Save(transaction *Transaction) error {
  args := m.Called(transaction)
  return args.Error(0)
}

func (m *mockTransactionRepository) GetAll() ([]Transaction, error) {
  args := m.Called()
  return args.Get(0).([]Transaction), args.Error(1)
}

func TestTransactionService(t *testing.T) {
	transactionRepository := new(mockTransactionRepository)
	transactionService := NewTransactionService(transactionRepository)

	t.Run("Should create a new transaction", func(t *testing.T) {	
		transactionRepository.On("Save", mock.Anything).Return(nil)
		input := CreateTransactionInput{
			Amount: 123,
			Description: "teste",
			TransactionType: Credit,
		}
		output, err := transactionService.CreateTransaction(input)
		transactionRepository.AssertExpectations(t)
		assert.Nil(t, err)
		assert.NotNil(t, output.ID) 
		assert.Equal(t, 123, output.Amount) 
		assert.Equal(t, "teste", output.Description) 
	})

	t.Run("Should list all transactions", func(t *testing.T) {
		createdAt, _ := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")
		transactionRepository.On("GetAll").Return([]Transaction{
			{
				ID: "123",
				Amount: 123,
				Description: "teste",
				TransactionType: Credit,
				CreatedAt: createdAt,
			},
		}, nil)
		output, err := transactionService.GetAllTransactions()
		transactionRepository.AssertExpectations(t)
		assert.Nil(t, err)
		assert.NotNil(t, output)
		assert.GreaterOrEqual(t, len(output), 0)
	})
}