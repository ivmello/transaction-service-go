package transaction

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type handler struct{
	service TransactionService
}

func NewTransactionHandler(service TransactionService) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) CreateTransaction(ctx *fiber.Ctx) error {
	var errors []*IError
	var body CreateTransactionInput
	if err := ctx.BodyParser(&body); err != nil {
		return err
	}
	var Validator = validator.New()
	err := Validator.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el IError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			errors = append(errors, &el)
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}
	transaction, err := h.service.CreateTransaction(body)
	if err != nil {
		return err
	}
	return ctx.JSON(transaction)
}

func (h *handler) GetAllTransactions(ctx *fiber.Ctx) error {
	transactions, err := h.service.GetAllTransactions()
	if err != nil {
		return err
	}
	return ctx.JSON(transactions)
}