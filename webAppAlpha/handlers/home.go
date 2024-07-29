package handlers

import (
	"fmt"
	"net/http"
)

// Home handles requests to the root URL.
func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorPage(w, r)
		return
	}
	fmt.Fprintf(w, "Welcome to the HomePage!")
}
