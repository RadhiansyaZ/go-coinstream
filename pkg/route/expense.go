package route

import (
	"github.com/go-chi/chi/v5"
	"go-coinstream/pkg/handler"
)

func ExpenseRouter(routeHandler *handler.Handlers) *chi.Mux {
	route := chi.NewRouter()

	route.Get("/", routeHandler.GetAllExpenses)
	route.Post("/", routeHandler.CreateExpense)
	route.Get("/{id}", routeHandler.GetExpenseByID)
	route.Put("/{id}", routeHandler.UpdateExpense)
	route.Delete("/{id}", routeHandler.DeleteExpenseByID)

	return route
}
