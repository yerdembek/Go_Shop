package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Connect() {
	fmt.Println("Start of work...")

	connStr := "user=postgres password=123AAss host=localhost dbname=web1 sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Ошибка при открытии соединения:", err)
		return
	}
	err = db.Ping()

	if err != nil {
		fmt.Println("Ошибка при проверке соединения:", err)
		return
	}
	query, err := db.Query("DELETE FROM users")
	fmt.Println(query)
	if err != nil {
		return
	}

	fmt.Println("Successfully connected!")
}
