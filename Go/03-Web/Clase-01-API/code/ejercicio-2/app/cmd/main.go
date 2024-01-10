// This script contains the solution for the task 2 (Create a simple REST API using the Body)
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/fatih/color"
)

var (
	green = color.New(color.FgGreen).SprintFunc()
)

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func main() {

	fmt.Println(green("Servidor iniciado Correctamente"))
	// Creation handler for the  endpoint greetings
	// Endpoint created to get a POST request and return a greeting message
	greetings := func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		// Get the body of the request and return a greeting message
		case "POST":
			// Read the body of the request

			var p Person // Creation of the struct to use the package json

			// Decode the body of the request in a Person struct
			err := json.NewDecoder(r.Body).Decode(&p)
			if err != nil {
				http.Error(w, "Error reading the body", http.StatusInternalServerError)
				return
			}
			// Return a greeting message using the first and last name of the person
			fmt.Fprintf(w, "Hello %s %s", p.FirstName, p.LastName)
		// Return an error if the method is not allowed
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}

	// Register the endpoint greetings

	http.HandleFunc("/greetings", greetings)

	// Start the server

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

	fmt.Println(green("Servidor apagado"))
}
