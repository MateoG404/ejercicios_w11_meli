package main

import (
	"app_ejercicio4/app_ejercicio5/internal"
	"fmt"
)

func main() {

	fmt.Println("Inicio de programa")
	var animalInput string
	fmt.Println("Ingrese el animal que desea alimentar")
	fmt.Scanln(&animalInput)
	animalFunc := internal.Orchestator(animalInput) // Use the imported package and function
	if animalFunc != nil {
		fmt.Println(animalFunc())
	}
}
