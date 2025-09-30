package app

import (
	"fmt"
	"net/http"
)

func Run() {
	fs := http.FileServer(http.Dir("./web/static"))
	http.Handle("/web/static/", http.StripPrefix("/web/static/", fs))

	http.HandleFunc("/", HomePage)
	http.HandleFunc("/login", LoginPage)
	http.HandleFunc("/register", RegisterPage)

	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe("localhost:8080", nil)
}
