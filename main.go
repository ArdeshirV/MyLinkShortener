package main

import (
	"fmt"
	"log/slog"
	"math/rand"
	"net/http"
)

func main() {
	fmt.Println(Prompt("My Link Shortener"))

	http.HandleFunc("/", redirectHandler)
	http.HandleFunc("/shorten", shortenHandler)

	slog.Info("Server started on http://localhost:8080")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		slog.Error(err.Error())
	}
}

const (
	baseURL     = "http://localhost:8080"
	shortURLLen = 4
)

var (
	urlMap  = make(map[string]string)
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[1:]
	if url, ok := urlMap[shortURL]; ok {
		http.Redirect(w, r, url, http.StatusFound)
	} else {
		http.NotFound(w, r)
	}
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	url := r.FromValue("url")
	if url == "" {
		http.Error(w, "URL cannot be empty", http.StatusBadRequest)
		return
	}

	shortURL := generateShortURL()
	urlMap[shortURL] = url

	shortenedURL := baseURL + shortURL
	fmt.Fprintf(w, "Shortened URL: %s", shortenedURL)
}

func generateShortURL() string {
	shortURL := make([]rune, shortURLLen)
	for i := range shortURL {
		shortURL[i] = letters[rand.Intn(len(letters))]
	}
	return string(shortURL)
}
