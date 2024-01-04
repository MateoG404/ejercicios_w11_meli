package main

import (
	"errors" // Import the "errors" package
	"fmt"
)

// Create a new type of error
type MyError struct {
	msg   string
	cause string
}

// Method
// Error() string
func (e *MyError) Error() string {

	return e.msg
}

func main() {
	fmt.Println("Inicio de programa")

	// Package errors
	// Method is

	// Creation first error
	err := errors.New("Error 1")
	// Creation wrap error
	wrapErr := fmt.Errorf("Error 2 %w", err)

	if errors.Is(wrapErr, err) {
		fmt.Println("Error 1 is wraped in Error 2")
	} else {
		fmt.Println("Error 1 is not wraped in Error 2")
	}

	// Errors using As()

	// Creation first error

	errorAs := MyError{msg: "Error 3", cause: "Base de datos no disponible"}
	// Creation wrap error
	errorWrapAs := fmt.Errorf("Error 4 %w", &errorAs)

	// Use As() to get the error, message and cause
	var err2 *MyError
	if errors.As(errorWrapAs, &err2) {
		fmt.Println("Error 3 is wraped in Error 4")
		fmt.Println("Error 3 message:", err2.msg, "Error 3 cause:", err2.cause)
	}
	fmt.Println("Fin de programa")
}
