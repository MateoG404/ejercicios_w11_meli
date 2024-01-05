// Script created to solve the excerise 1 of the class Panic & Defer in Go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Inicio del programa para leer datos del estudio contable ")

	// Read the file
	file, err := os.Open("./customers.txt")

	// Defer if the file doesnt exist
	defer fmt.Println("Ejecucion Finalizada")

	// Check if the file exists if not trough a panic
	if err != nil {
		panic("The indicated file was not found or is damaged")
	}

	// Read the file if the file exists
	scanner := bufio.NewScanner(file)

	// Create a variable to store the customers information
	var customers []string

	// Read all the lines of the file

	for scanner.Scan() {
		customers = append(customers, scanner.Text())

	}
	// Show the customers information
	fmt.Println("Customers: ")
	fmt.Println(customers)
	// Close the file
	file.Close()

}
