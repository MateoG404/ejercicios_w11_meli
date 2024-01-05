// This is the code to solve the Excercise 2 section Errors of the Go course
package main

import (
	"errors"
	"fmt"
)

type ErrorMain struct {
	msg string
}

// Method Error() string to return the error message

func (e *ErrorMain) Error() string {
	return e.msg
}

// Method NewErrorMain() to create a new error
func NewErrorMain() error {
	return &ErrorMain{msg: "Error: salary is less than 10000"}
}
func main() {
	var salary int
	fmt.Println("Ingrese su salario: ")
	fmt.Scan(&salary)

	// Creation error to be wrapped
	err := NewErrorMain()

	// Condition to validate if the salary is less than 10000
	if salary < 10000 {
		// Creation error
		errorWrap := fmt.Errorf("Error: Created by  %w", err)
		if errors.Is(errorWrap, err) {
			fmt.Println(" Error wrapped by Error main")
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println("El salario es mayor o igual a 10000")
	}

}
