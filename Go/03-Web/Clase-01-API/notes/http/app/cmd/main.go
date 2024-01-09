// This scripts contains the way to create a server as a native way using the net/http package

package main

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
)

var (
	green = color.New(color.FgGreen).SprintFunc()
	blue  = color.New(color.FgBlue).SprintFunc()
)

func main() {
	fmt.Println(green("Servidor iniciado"))

	// 1. Create the cases  for the handlerTest

	handlerTest := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(blue("Saludos desde el endpoint prueba"))
		// Escribe en el writer para retornarlo en el response
		fmt.Fprintf(w, "Saludos desde el endpoint prueba")
		// Start the server
	}

	handlerCases := func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			fmt.Fprintf(w, "Saludos desde el endpoint prueba usando GET")
		case "POST":
			fmt.Fprintf(w, "Saludos desde el endpoint prueba usando POST")

		}
	}

	// _________________________________________________________

	// 2. Register the funcionalities for the handlerTest
	http.HandleFunc("/prueba", handlerTest)

	// 2. Register a second handler
	http.HandleFunc("/casos", handlerCases)

	// _________________________________________________________

	// 3. Start the server
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println(err)
	}
}
