// File: api/index.go
package main

import (
	"fmt"
	"net/http"
)

// Handler is automatically recognized by Vercel as the entry point
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Paisa finance tracker is running on Vercel!")
}
