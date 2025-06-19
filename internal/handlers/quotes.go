package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"test/internal/models"
	"test/internal/storage"
)

type QuoteHandler struct {
	store *storage.MemoryStore
}

func NewQuoteHandler(store *storage.MemoryStore) *QuoteHandler {
	return &QuoteHandler{store: store}
}

func (h *QuoteHandler) AddQuote(w http.ResponseWriter, r *http.Request) {
	var q models.Quote
	if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	q = h.store.Add(q)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(q)
}

func (h *QuoteHandler) GetQuotes(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")
	var quotes []models.Quote

	if author != "" {
		quotes = h.store.GetByAuthor(author)
	} else {
		quotes = h.store.GetAll()
	}

	json.NewEncoder(w).Encode(quotes)
}

func (h *QuoteHandler) GetRandomQuote(w http.ResponseWriter, r *http.Request) {
	q, err := h.store.GetRandom()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(q)
}

func (h *QuoteHandler) DeleteQuote(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Ошибка формата", http.StatusBadRequest)
		return
	}

	if err := h.store.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
