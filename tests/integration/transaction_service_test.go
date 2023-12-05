package integration

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
	"transaction-service-go/internal/app/transaction"
	"transaction-service-go/internal/infra/database"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestTransactionService_CreateTransaction(t *testing.T) {
	err := godotenv.Load("../../.env.testing")
  if err != nil {
    log.Fatal("Error loading .env file")
  }
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlPort := os.Getenv("MYSQL_PORT")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlDatabase)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	mysqlAdapter := database.NewMysqlAdapter(db)
	transactionService := transaction.NewTransactionService(mysqlAdapter)

	t.Run("Should create a new transaction", func(t *testing.T) {
		input := transaction.CreateTransactionInput{
			Amount: 123,
			Description: "teste",
			TransactionType: transaction.Credit,
		}
		output, err := transactionService.CreateTransaction(input)
		assert.Nil(t, err)
		assert.NotNil(t, output.ID)
		assert.Equal(t, 123, output.Amount)
		assert.Equal(t, "teste", output.Description)
	})
	
	t.Run("Should list all transactions", func(t *testing.T) {
		input := transaction.CreateTransactionInput{
			Amount: 123,
			Description: "teste",
			TransactionType: transaction.Credit,
		}
		_, _ = transactionService.CreateTransaction(input)
		output, err := transactionService.GetAllTransactions()
		assert.Nil(t, err)
		assert.NotNil(t, output)
		assert.GreaterOrEqual(t, len(output), 1)
	})
}