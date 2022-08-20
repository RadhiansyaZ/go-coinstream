package route

import (
	"github.com/go-chi/chi/v5"
	"go-coinstream/pkg/handler"
)

func UserRouter(routeHandler handler.UserHandler) *chi.Mux {
	route := chi.NewRouter()

	route.Post("/register", routeHandler.Register)
	route.Post("/login", routeHandler.Login)

	return route
}
