package main

import (
	"fmt"
	"log/slog"
	"math/rand"
	"net/http"
)

func main() {
	fmt.Println(Prompt("My Link Shortener"))
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
