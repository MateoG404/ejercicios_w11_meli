// Script created to solve the exercise 3 of the class 6 Panic & Defer of the Go course

package main

import (
	"ejercicio_3/app_ejercicio3/internal"
	"errors"
	"fmt"

	"github.com/fatih/color"
)

var (
	red = color.New(color.FgRed).SprintFunc()
)

func recoverPanic() {
	if r := recover(); r != nil {
		fmt.Println("RECOVERED by the main", r)
	}
}

func verifyValuesInput(fileName, nameClient, idClient string, phoneNumberClient, homeNumberClient int) (bool, error) {
	if fileName == "" || nameClient == "" || idClient == "" || phoneNumberClient == 0 || homeNumberClient == 0 {
		return false, errors.New(red("error: Invalid input values"))
	}
	return true, nil
}
func main() {

	// Defers
	defer fmt.Println("End of execution")
	defer func() {
		if internal.PanicErros {
			fmt.Println(errors.New(red("Several errors were detected at runtime")))
		}
	}()
	defer func() {
		r := recover()
		if r != nil {
			internal.PanicErros = true
			fmt.Println("aca")
			fmt.Println("RECOVERED by the main", r)
		}
	}()

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

	// Verify if the values of the new client are valid
	if ok, err := verifyValuesInput(fileName, nameClient, idClient, phoneNumberClient, homeNumberClient); !ok {
		fmt.Println(err)
		return
	}
	// Verify if the file exists
	_, exists := newClient.VerifyFile()

	if !exists {
		panic(errors.New(red("error: File does not exist")))
	}
	// Verify if the new client exists
	if newClient.VerifyClientFile() { // If the client exists, the program will panic
		fmt.Println("Client already existssssss")
		panic(errors.New(red("error: Client already exists")))
	}

}
