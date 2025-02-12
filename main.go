package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AkiGR/go-shortly/handlers"
	"github.com/AkiGR/go-shortly/storage"
)

func main() {
	storage.InitDB()

	http.HandleFunc("/shorten", handlers.ShortenURL)

	fmt.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
