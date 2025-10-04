package app

import (
	"First_Structured_GO_Project/internal/database"
	"First_Structured_GO_Project/web/handlers"
	"fmt"
)

func Run() {
	fmt.Println("In Run function")

	database.Connect()

	handlers.Handle()
}
