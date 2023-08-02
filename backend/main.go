package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Book структура представляет информацию о книге
type Book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Genre    string `json:"genre"`
	Price    float64 `json:"price"`
	Quantity int    `json:"quantity"`
}

var books []Book

func main() {
	// Добавим несколько книг для примера
	books = []Book{
		{ID: 1, Title: "Book 1", Author: "Author 1", Genre: "Fiction", Price: 29.99, Quantity: 10},
		{ID: 2, Title: "Book 2", Author: "Author 2", Genre: "Non-Fiction", Price: 19.99, Quantity: 5},
	}

	http.HandleFunc("/books", getBooksHandler)
	http.HandleFunc("/books/add", addBookHandler)

	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getBooksHandler(w http.ResponseWriter, r *http.Request) {
	// Отправляем список книг в формате JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func addBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Получаем данные из тела запроса
	var newBook Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Генерируем новый ID для книги
	newBook.ID = len(books) + 1

	// Добавляем книгу в список
	books = append(books, newBook)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}
