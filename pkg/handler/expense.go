package handler

import (
	"github.com/gofiber/fiber/v2"
	"go-coinstream/pkg/dto"
	"go-coinstream/pkg/service"
)

type ExpenseHandlers interface {
	GetAllExpenses(ctx *fiber.Ctx) error
	GetExpenseByID(ctx *fiber.Ctx) error
	CreateExpense(ctx *fiber.Ctx) error
	UpdateExpense(ctx *fiber.Ctx) error
	DeleteExpenseByID(ctx *fiber.Ctx) error
}

type handlers struct {
	service *service.ExpenseService
}

func NewHttpExpenseHandler(expenseService service.ExpenseService) *handlers {
	return &handlers{
		service: &expenseService,
	}
}

func (h *handlers) GetAllExpenses(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")

	expenses := (*h.service).FindAll()

	return ctx.JSON(expenses)
}
func (h *handlers) GetExpenseByID(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")
	id := ctx.Params("id")

	expense, err := (*h.service).FindById(id)
	if err != nil {
		return err
	}

	return ctx.JSON(expense)
}
func (h *handlers) CreateExpense(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")
	exp := new(dto.ExpenseRequest)
	if err := ctx.BodyParser(exp); err != nil {
		return err
	}

	expense, err := (*h.service).Add(exp)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(expense)
}

func (h *handlers) UpdateExpense(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")
	id := ctx.Params("id")

	exp := new(dto.ExpenseRequest)
	if err := ctx.BodyParser(exp); err != nil {
		return err
	}

	expense, err := (*h.service).Update(id, exp)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusAccepted).JSON(expense)
}
func (h *handlers) DeleteExpenseByID(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")
	id := ctx.Params("id")

	err := (*h.service).Delete(id)
	if err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
