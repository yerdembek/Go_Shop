package handlers

import (
	"fmt"
	"net/http"
)

func Handle() {
	fs := http.FileServer(http.Dir("./web/static"))
	http.Handle("/web/static/", http.StripPrefix("/web/static/", fs))

	http.HandleFunc("/", HomePage)
	http.HandleFunc("/login", LoginPage)
	http.HandleFunc("/register", RegisterPage)
	http.HandleFunc("/posting", PostAdPage)

	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe("localhost:8080", nil)
}
