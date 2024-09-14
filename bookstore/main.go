package main

import (
	"bookstore/database"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Book struct {
	ID    uint    `json:"id"`
	Title string  `json:"title"`
	Price float32 `json:"price"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type Response struct {
	Message string `json:"message"`
}

func main() {
	// load enviroment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load enviroment variables")
	}

	var port string
	port = os.Getenv("PORT")
	if port == "" {
		port = "5050"
	}

	http.HandleFunc("POST /newbook", createBook)

	fmt.Println("Server running on port" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	db := database.DataBase()

	err := json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{
			Message: "Server Error",
		})
	}
	fmt.Println(book)

	//save book to db
	result, err := db.Exec("INSERT INTO books (id,title,price) VALUES($1,$2,$3)", book.ID, book.Title, book.Price)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{
			Message: "Something went wrong",
		})
	}

	fmt.Println(result)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		Message: "Book Created Successfully",
	})
}
