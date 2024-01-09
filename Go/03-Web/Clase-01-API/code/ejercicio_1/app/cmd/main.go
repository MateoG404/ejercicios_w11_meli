// This script contains the solution for the first task of the class API
// The task is to create a server using the net/http package with an endpoint that returns a text
package main

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
)

var (
	green = color.New(color.FgGreen).SprintFunc()
)

func main() {
	fmt.Println(green("Inicio del servidor"))

	// Define the endpoints in the handler
	ping := func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			fmt.Fprintf(w, "pong")
		default:
			fmt.Fprintf(w, "Error 405: Method not allowed")

		}
	}

	// Register the endpoints into the handler
	http.HandleFunc("/ping", ping)

	// Start the server

	// listen and serve
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(green("Finalizacion del servidor"))

}
