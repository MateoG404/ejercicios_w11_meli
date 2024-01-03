package main

import (
	"ejercicio_2/app_ejercicio2/internal"
	"fmt"
)

func main() {

	// Create a new Person
	person := internal.Person{ID: 1, Name: "Mateo Gutierrez", DateOfBirth: "01/01/2000"}

	// Create a new Employee
	employee := internal.Employee{ID: 1, Position: "Software Developer", Person: person}
	// Print the employee
	fmt.Println("Nuestro nuevo empleado es :")
	fmt.Println()
	employee.PrintEmployee()
}
