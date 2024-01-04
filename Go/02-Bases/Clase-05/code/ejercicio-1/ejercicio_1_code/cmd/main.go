// Code created by: @mateosky
// This code solves the exercise 1 of the class 5 (interfaces) of the Go course
package main

import (
	"ejercicio_1_code/ejercicio_1_code/internal"
	"fmt"
)

func main() {
	// Variables
	var name string
	var apellido string
	var dni int
	var fecha string

	// Input data
	fmt.Println("Hola, por favor ingresa tu nombre")

	fmt.Scanln(&name)

	fmt.Println("Por favor ingresa tu apellido")
	fmt.Scanln(&apellido)

	fmt.Println("Por favor ingresa tu DNI")
	fmt.Scanln(&dni)

	fmt.Println("Por favor ingresa tu fecha de nacimiento")
	fmt.Scanln(&fecha)

	// Create the student
	student := internal.Student{Name: name, Apellido: apellido, DNI: dni, FechaNacimiento: fecha}

	// Print the student
	fmt.Println("Tu informacion es la siguiente:")
	student.PrintStudent()
}
