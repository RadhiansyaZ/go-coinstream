package main

import (
	"github.com/joho/godotenv"
	app "go-coinstream/cmd"
)

func main() {
	godotenv.Load()
	app.Run()
}
