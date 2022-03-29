package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type API struct{}

var books = []string{"book 1", "book 2", "book 3"}

func (a *API) getBooks(w http.ResponseWriter, r *http.Request) {
	limitParam := r.URL.Query().Get("limit")

	limit, err := strconv.Atoi(limitParam)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if limit < 1 || limit > len(books) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(books[:limit])
}

func (a *API) getBook(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if id < 1 || id > len(books) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(books[id])
}
