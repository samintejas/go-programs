package main

import (
	"animal-webapp/handlers" // Adjust the import path to match your module
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

const PORT string = "8081"

func main() {

	log := slog.New(slog.NewTextHandler(os.Stderr, nil))

	log.Info("Starting server!")
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/animal", handlers.AddAnimal)
	log.Info("Server started! ", "Port", PORT)

	err := http.ListenAndServe(":"+PORT, nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
