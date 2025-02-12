package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/AkiGR/go-shortly/storage"
)

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	originalURL := r.URL.Query().Get("url")

	if originalURL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	shortID := generateShortID()

	db := storage.GetDB()
	stmt, err := db.Prepare("INSERT INTO urls(original_url, short_id) VALUES(?, ?)")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	_, err = stmt.Exec(originalURL, shortID)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Shortened URL: http://localhost:8080/%s", shortID)
}

func generateShortID() string {
	rand.Seed(time.Now().UnixNano())
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	shortID := make([]byte, 6)
	for i := range shortID {
		shortID[i] = letters[rand.Intn(len(letters))]
	}
	return string(shortID)
}

