// Script created to solve the exercise 3 of the class 6 Panic & Defer of the Go course

package main

import (
	"ejercicio_3/app_ejercicio3/internal"
	"errors"
	"fmt"
)

func recoverPanic() {
	if r := recover(); r != nil {
		fmt.Println("RECOVERED by the main", r)
	}
}
func main() {

	// Defers
	defer recoverPanic()
	// Verify if the new client exists
	var fileName string
	var nameClient string
	var idClient string
	var phoneNumberClient int
	var homeNumberClient int

	// Store the values of the new possible client
	fmt.Println("Ingresa el nombre del archivo donde deseas verificar la existencia de tu nuevo posible cliente")
	fmt.Scanln(&fileName)
	fmt.Println("Hola, ingresa el nombre del cliente que deseas verificar")
	fmt.Scanln(&nameClient)
	fmt.Println("Ingresa el ID del cliente que deseas verificar")
	fmt.Scanln(&idClient)
	fmt.Println("Ingresa el número de teléfono del cliente que deseas verificar")
	fmt.Scanln(&phoneNumberClient)
	fmt.Println("Ingresa el número de casa del cliente que deseas verificar")
	fmt.Scanln(&homeNumberClient)

	// Create the new possible client
	newClient := internal.NewClient(fileName, nameClient, idClient, phoneNumberClient, homeNumberClient)

	// Verify if the new client exists
	if newClient.VerifyClientFile() { // If the client exists, the program will panic
		fmt.Println("El cliente ya existe")
		panic(errors.New("error: Client already exists"))
	}

}
