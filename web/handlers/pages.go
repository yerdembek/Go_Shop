package handlers

import (
	"net/http"
	"os"
)

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	htmlFile, err := os.ReadFile("web/templates/register.html")
	if err != nil {
		http.Error(w, "Ошибка чтения файла: "+err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(htmlFile)
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	htmlFile, err := os.ReadFile("web/templates/login.html")
	if err != nil {
		http.Error(w, "Ошибка чтения файла: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(htmlFile)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	htmlFile, err := os.ReadFile("web/templates/index.html")
	if err != nil {
		http.Error(w, "Ошибка чтения файла: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(htmlFile)
}

func PostAdPage(w http.ResponseWriter, r *http.Request) {
	htmlFile, err := os.ReadFile("web/templates/post_ad.html")
	if err != nil {
		http.Error(w, "Ошибка чтения файла: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(htmlFile)
}
