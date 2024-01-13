package main

import (
	"log"
	"transaction-service-go/internal/app/transaction"
	"transaction-service-go/internal/infra/database"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
	app := fiber.New()
	db, err := database.NewPostgresConnection()
	if err != nil {
		log.Fatal(err)
	}
	postgresAdapter := database.NewPostgresAdapter(db)
	transaction.NewTransactionRoutes(app, postgresAdapter)
	app.Listen(":5000")
}
