package route

import (
	"github.com/go-chi/chi/v5"
	"go-coinstream/pkg/handler"
)

func IncomeRouter(routeHandler handler.IncomeHandler) *chi.Mux {
	route := chi.NewRouter()

	route.Get("/", routeHandler.GetAllIncomes)
	route.Post("/", routeHandler.CreateIncome)
	route.Get("/{id}", routeHandler.GetIncomeByID)
	route.Put("/{id}", routeHandler.UpdateIncome)
	route.Delete("/{id}", routeHandler.DeleteIncomeByID)

	return route
}
