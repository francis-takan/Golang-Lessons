package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	DateOfBirth string `json:"date_of_birth"`
	ImageUrl    string `json:"image_url"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type Response struct {
	Message string `json:"message"`
}

func main() {
	const port = ":8000"
	mux := http.NewServeMux()

	/**
	A route that gets user details
	*/
	mux.HandleFunc("GET /user", getUser)
	mux.HandleFunc("POST /register", createUser)

	fmt.Printf("Server running on port.... %v\n", port)
	log.Fatal(http.ListenAndServe(port, mux))
}

func getUser(w http.ResponseWriter, r *http.Request) {
	suzan := User{
		FirstName:   "Suzan",
		LastName:    "Kay",
		PhoneNumber: "+256 777 29 6895",
		DateOfBirth: "20-February-1999",
		ImageUrl:    "https://i.ibb.co/rvZnMYL/profilewoman.jpg",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(suzan)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User

	er := ErrorResponse {
		Message: "Server error",
	}

	res := ErrorResponse {
		Message: "Application successful",
	}
	
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		json.NewEncoder(w).Encode(er)
		return
	}

	fmt.Println(user)

	json.NewEncoder(w).Encode(res)
}


