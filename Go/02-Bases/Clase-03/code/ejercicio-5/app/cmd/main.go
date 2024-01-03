package main

import (
	"app/internal"
	"fmt"
)

func main() {

	fmt.Println("Inicio de programa")
	var animalInput string
	fmt.Println("Ingrese el animal que desea alimentar")
	fmt.Scanln(&animalInput)
	animalFunc := internal.Orchestator(animalInput)
	if animalFunc != nil {
		fmt.Println(animalFunc())
	}
}
