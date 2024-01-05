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
	var salary int
	fmt.Println("Ingrese su salario: ")

	fmt.Scan(&salary)
	if salary < 15000 {
		var err Error
		fmt.Println(err.Error())
	} else {
		fmt.Println("Must pay tax")
	}

}
