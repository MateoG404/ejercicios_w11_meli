// This script is created to test the chi package

package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {

	// Create a new router to append the new endpoints

	router := chi.NewRouter()

	// Creation new handler for a specific method get

	router.Get("/prueba", func(w http.ResponseWriter, r *http.Request) {
		// Set code for the response
		w.WriteHeader(200)

		w.Write([]byte("Hola Mateo"))
	})

	// Start the server

	http.ListenAndServe(":3000", router)
}
