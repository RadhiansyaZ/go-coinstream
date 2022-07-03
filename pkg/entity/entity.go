package entity

type Expense struct {
	ID       string
	Name     string
	Amount   float64
	Category string
	Date     string // yyyy-mm-dd
}

type Income struct {
	ID     string
	Name   string
	Amount float64
	Date   string // yyyy-mm-dd
}
