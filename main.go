package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func open_file(filename string) {
	fmt.Println("Opening file: " + filename)
}

func getFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}

type LoginDetails struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var details LoginDetails

	err := json.NewDecoder(r.Body).Decode(&details)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if details.Username != "admin" || details.Password != "admin" {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	} else {
		fmt.Fprintf(w, "Login successful")
	}
}

func main() {
	http.HandleFunc("/login", loginHandler)
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
