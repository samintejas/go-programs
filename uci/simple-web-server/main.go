package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {

	log := slog.New(slog.NewTextHandler(os.Stdout, nil))

	log.Info("Defining routes ..")
	r := mux.NewRouter()

	r.HandleFunc("/products", getProducts).Methods("GET")
	r.HandleFunc("/product/{id:[0-9]+}", getProduct).Methods("GET")
	r.HandleFunc("/products", addProduct).Methods("POST")
	r.HandleFunc("/product/{id:[0-9]+}", updateProduct).Methods("PUT")
	r.HandleFunc("/product/{id:[0-9]+}", deleteProduct).Methods("DELETE")

	log.Info("Starting server on port 8080 ..")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Error("Could not start server .. Exiting .. ")
	}

}

type Product struct {
	ID   int64
	Name string
	Type string
}

var products = make(map[int64]Product)

func getProducts(w http.ResponseWriter, r *http.Request) {

	var productList []Product
	for _, p := range products {
		productList = append(productList, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(productList)

}
func getProduct(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Not a numerical ID", http.StatusNotFound)
		return
	}

	p, excist := products[idNum]
	if !excist {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)

}
func addProduct(w http.ResponseWriter, r *http.Request) {
	var p Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if p.ID == 0 || p.Name == "" || p.Type == "" {
		http.Error(w, "Invalid product data", http.StatusBadRequest)
		return

	}
	products[p.ID] = p
	w.WriteHeader(http.StatusCreated)

}
func updateProduct(w http.ResponseWriter, r *http.Request) {}
func deleteProduct(w http.ResponseWriter, r *http.Request) {}
