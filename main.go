package main

import (
	"net/http"

	"github.com/bookcrud/handlers"
)

func main() {
	http.HandleFunc("/books", handlers.GetBooks)

	http.HandleFunc("/book", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handlers.CreateBook(w, r)
		case http.MethodGet:
			handlers.GetBook(w, r)
		case http.MethodPut:
			handlers.UpdateBook(w, r)
		case http.MethodDelete:
			handlers.DeleteBook(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}
