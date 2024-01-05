package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Inicio del programa para leer daros del estudio contable ")
	// Defer if the file doesnt exist
	defer fmt.Println("Ejecucion Finalizada")
	// Read the file
	file, err := os.Open("./customers.txt")

	if err != nil {
		panic(err)
	}

	// Read the file
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
