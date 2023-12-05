package transaction

import (
	"github.com/gofiber/fiber/v2"
)

func NewTransactionRoutes(app *fiber.App, transactionRepository TransactionRepository) {
	transactionService := NewTransactionService(transactionRepository)
	handler := NewTransactionHandler(transactionService)
	api := app.Group("/transactions")
	api.Post("/", handler.CreateTransaction)
	api.Get("/", handler.GetAllTransactions)
}