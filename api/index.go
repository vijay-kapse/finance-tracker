package main

import (
    "fmt"
    "net/http"
)

// Handler is the entry point for Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Paisa finance tracker is running on Vercel!")
}

