package handler

import (
	"fmt"
	"net/http"
)

// Handler function required by Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Paisa finance tracker is running on Vercel!")
}
