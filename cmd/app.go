package cmd

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"go-coinstream/pkg/dto"
	"go-coinstream/pkg/entity"
	"go-coinstream/pkg/handler"
	"go-coinstream/pkg/repository"
	"go-coinstream/pkg/service"
	"log"
	"net/http"
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

func handlerTest4(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := createConnection()
	if err != nil {
		panic(err)
	}
	expenseRepo := repository.NewExpenseRepository(db)

	vars := mux.Vars(r)

	err1 := expenseRepo.Delete(vars["id"])
	if err1 != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error":"Expense Not Found"}`))
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func handlerTest3(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := createConnection()
	if err != nil {
		panic(err)
	}
	expenseRepo := repository.NewExpenseRepository(db)

	vars := mux.Vars(r)

	expResult, err := expenseRepo.FindById(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error":"Expense Not Found"}`))
		return
	}

	w.WriteHeader(http.StatusAccepted)
	result, _ := json.Marshal(expResult)
	w.Write(result)
}

func handlerTest2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := createConnection()
	if err != nil {
		panic(err)
	}

	var expenseReq dto.ExpenseRequest
	err1 := json.NewDecoder(r.Body).Decode(&expenseReq)
	if err1 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"Error unmarshalling"}`))
		return
	}

	expenseRepo := repository.NewExpenseRepository(db)

	vars := mux.Vars(r)

	expense := entity.Expense{
		Name:     expenseReq.Name,
		Amount:   expenseReq.Amount,
		Category: expenseReq.Category,
		Date:     expenseReq.Date,
	}

	expResult, err := expenseRepo.Update(vars["id"], &expense)

	w.WriteHeader(http.StatusAccepted)
	result, err := json.Marshal(expResult)
	w.Write(result)
}

func handlerTest1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := createConnection()
	if err != nil {
		panic(err)
	}

	var expenseReq dto.ExpenseRequest
	err1 := json.NewDecoder(r.Body).Decode(&expenseReq)

	if err1 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"Error unmarshalling"}`))
		return
	}

	expenseRepo := repository.NewExpenseRepository(db)

	expense := entity.Expense{
		Name:     expenseReq.Name,
		Amount:   expenseReq.Amount,
		Category: expenseReq.Category,
		Date:     expenseReq.Date,
	}

	expResult, err := expenseRepo.Add(&expense)

	w.WriteHeader(http.StatusAccepted)
	result, err := json.Marshal(expResult)
	w.Write(result)
}

func handlerTest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := createConnection()
	if err != nil {
		panic(err)
	}

	expenseRepo := repository.NewExpenseRepository(db)

	expenses, err := expenseRepo.FindAll()

	result, err := json.Marshal(expenses)

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
