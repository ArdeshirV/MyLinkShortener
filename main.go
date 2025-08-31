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
    baseURL = "http://localhost:8080"
    shortURLLen = 4
)

var (
    urlMap = make(map[string]string)
    letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)


