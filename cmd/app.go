package cmd

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"go-coinstream/pkg/handler"
	"go-coinstream/pkg/repository"
	"go-coinstream/pkg/route"
	"go-coinstream/pkg/service"
	"log"
	"net/http"
	"os"
)

func createConnection() (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/coinstream?sslmode=disable",
		os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"))

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Successfully connected to Postgres!")
	return db, nil
}

func Run() {
	db, err := createConnection()

	if err != nil {
		panic(err)
	}

	expenseRepo := repository.NewExpenseRepository(db)
	expenseService := service.NewExpenseService(expenseRepo)
	expenseHandler := handler.NewHttpExpenseHandler(expenseService)
	expenseRouter := route.ExpenseRouter(expenseHandler)

	incomeRepo := repository.NewIncomeRepository(db)
	incomeService := service.NewIncomeService(incomeRepo)
	incomeHandler := handler.NewHttpIncomeHandler(incomeService)
	incomeRouter := route.IncomeRouter(incomeHandler)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewHttpUserHandler(userService)
	userRouter := route.UserRouter(userHandler)

	r := chi.NewRouter()
	r.Mount("/expense", expenseRouter)
	r.Mount("/income", incomeRouter)
	r.Mount("/user", userRouter)

	log.Println("Server starting at PORT 8000")

	log.Fatal(http.ListenAndServe(":8000", r))
}
