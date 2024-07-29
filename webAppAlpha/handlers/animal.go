package handlers

import (
	"animal-webapp/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var animals = []Animal{
	{ID: 1, Name: "Buddy", Type: "Dog"},
	{ID: 2, Name: "Whiskers", Type: "Cat"},
}

// AddAnimal handles POST requests to add an animal.
func AddAnimal(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var animal models.Animal
	err = json.Unmarshal(body, &animal)
	if err != nil {
		http.Error(w, "Error unmarshaling JSON", http.StatusBadRequest)
		return
	}

	fmt.Printf("Received animal: %+v\n", animal)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Animal added successfully")
}

func getAnimals(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(animals)
}

func getAnimal(w http.ResponseWriter, r *http.Request) {
	// ... Retrieve animal by ID from database or in-memory storage
	// ... Encode and send response
}

func createAnimal(w http.ResponseWriter, r *http.Request) {
	// ... Decode request body to create an animal
	// ... Store the new animal in database or in-memory storage
	// ... Encode and send response
}

func updateAnimal(w http.ResponseWriter, r *http.Request) {
	// ... Retrieve animal by ID
	// ... Update animal data
	// ... Encode and send response
}

func deleteAnimal(w http.ResponseWriter, r *http.Request) {
	// ... Retrieve animal by ID
	// ... Delete animal from database or in-memory storage
	// ... Send success response
}
