package cmd

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"go-coinstream/pkg/handler"
	"go-coinstream/pkg/repository"
	"go-coinstream/pkg/service"
	"log"
)

func createConnection() (*sql.DB, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()

	connStr := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/coinstream?sslmode=disable",
		viper.Get("DATABASE_USERNAME"),
		viper.Get("DATABASE_PASSWORD"),
		viper.Get("DATABASE_HOST"),
		viper.Get("DATABASE_PORT"))

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

	app := fiber.New()

	exp := app.Group("/expense")
	exp.Get("/", expenseHandler.GetAllExpenses)
	exp.Post("/", expenseHandler.CreateExpense)
	exp.Get("/:id", expenseHandler.GetExpenseByID)
	exp.Put("/:id", expenseHandler.UpdateExpense)
	exp.Delete("/:id", expenseHandler.DeleteExpenseByID)

	log.Println("Server starting at PORT 8000")

	log.Fatal(app.Listen(":8000"))
}
