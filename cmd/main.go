package main

import (
	"log"
	"net/http"

	"test/internal/handlers"
	"test/internal/storage"

	"github.com/gorilla/mux"
)

func main() {
	store := storage.NewMemoryStore()
	handler := handlers.NewQuoteHandler(store)

	router := mux.NewRouter()

	router.HandleFunc("/quotes", handler.AddQuote).Methods("POST")
	router.HandleFunc("/quotes", handler.GetQuotes).Methods("GET")
	router.HandleFunc("/quotes/random", handler.GetRandomQuote).Methods("GET")
	router.HandleFunc("/quotes/{id}", handler.DeleteQuote).Methods("DELETE")

	log.Println("Сервер запущен http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
