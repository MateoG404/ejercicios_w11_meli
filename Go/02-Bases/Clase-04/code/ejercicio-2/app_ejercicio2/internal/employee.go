package internal

import "fmt"

// Employee is a struct that represents an employee
type Employee struct {
	ID       int
	Position string
	Person
}

// Methods
// PrintEmployee prints the employee

func (employee Employee) PrintEmployee() {
	fmt.Println("Employee ID: ", employee.ID)
	fmt.Println("Employee Position: ", employee.Position)
	fmt.Println("Employee Name: ", employee.Name)
	fmt.Println("Employee Date of Birth: ", employee.DateOfBirth)
}
