package handler

import (
	"fmt"
	"net/http"
)

// Handler is cool
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Go ONE!")
}
