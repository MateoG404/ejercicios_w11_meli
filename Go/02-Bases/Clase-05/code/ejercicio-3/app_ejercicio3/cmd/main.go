package main

import "fmt"

// Creation personalized error
type Error struct {
	msg string
}

// Method Error() string
func (e *Error) Error() string {
	e.msg = "Error:the salary entered does not reach the taxable minimum"
	return e.msg
}

func main() {
	// Declaration and initialization of variables
	var salary int
	fmt.Println("Ingrese su salario: ")

	// Read the salary from the keyboard
	fmt.Scan(&salary)

	if salary < 15000 {
		// Creation error
		var err Error
		fmt.Println(err.Error())
	} else {
		// Print message if the salary is greater than or equal to 15000
		fmt.Println("Must pay tax")
	}

}
