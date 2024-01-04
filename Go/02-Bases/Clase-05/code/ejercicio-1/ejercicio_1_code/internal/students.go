// This code is the implementation of the interface Student

package internal

import "fmt"

// Student is the struct that defines a student
type Student struct {
	Name            string
	Apellido        string
	DNI             int
	FechaNacimiento string
}

// Method

// PrintStudent prints the student
func (s *Student) PrintStudent() {
	fmt.Println("Name: ", s.Name)
	fmt.Println("Apellido: ", s.Apellido)
	fmt.Println("DNI: ", s.DNI)
	fmt.Println("Fecha: ", s.FechaNacimiento)
}
