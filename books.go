package main

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type BookHandler struct {
}

type Book struct {
	ID               string `json:"id"`
	Title            string `json:"title"`
	Author           string `json:"author"`
	PublishedDate    string `json:"published_date"`
	OriginalLanguage string `json:"original_language"`
}

var books = []*Book{
	{
		ID:               "1",
		Title:            "7 habits of Highly Effective People",
		Author:           "Stephen Covey",
		PublishedDate:    "15/08/1989",
		OriginalLanguage: "English",
	},
}

func listBooks() []*Book {
	return books
}

func (b BookHandler) ListBooks(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(listBooks())
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
func (b BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	book := getBook(id)
	if book == nil {
		http.Error(w, "Book not found", http.StatusNotFound)
	}
	err := json.NewEncoder(w).Encode(book)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
func (b BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	storeBook(book)
	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
func (b BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedBook := updateBook(id, book)
	if updatedBook == nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(updatedBook)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
func (b BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	found := deleteBook(id)
	if !found {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func getBook(id string) *Book {
	for _, book := range books {
		if book.ID == id {
			return book
		}
	}
	return nil
}

func storeBook(book Book) {
	books = append(books, &book)
}

func deleteBook(id string) bool {
	for i := range books {
		if books[i].ID == id {
			books = append(books[:i], books[i+1:]...)
			return true
		}
	}
	return false
}

func updateBook(id string, bookUpdate Book) *Book {
	for i, book := range books {
		if book.ID == id {
			books[i] = &bookUpdate
			return book
		}
	}
	return nil
}
