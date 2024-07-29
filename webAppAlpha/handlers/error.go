package handlers

import (
	"fmt"
	"net/http"
)

// ErrorPage handles 404 errors.
func ErrorPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Printf("%s : no such page", r.URL.Path)
	fmt.Fprintf(w, "404 - Page Not Found")
}
