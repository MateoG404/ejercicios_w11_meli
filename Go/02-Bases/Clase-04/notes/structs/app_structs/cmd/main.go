package main

import (
	"app_structs/app_structs/internal"
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Using structs")

	// Create a new person
	personA := internal.Person{FirstNmame: "Mateo", LastName: "Gutierrez", CellPhone: "123456789", Address: "Street 123", Password: "s"}
	personB := internal.Person{FirstNmame: "Camila", LastName: "Zambrano", CellPhone: "123456789", Password: "s"}
	fmt.Println(personA)
	fmt.Println(personB)

	// Store the personA in a JSON file
	personJSON, err := json.Marshal(personA)

	// Store the personB in a JSON file
	personJSONB, err := json.Marshal(personB)
	if err != nil {
		fmt.Println("Error marshaling person:", err)
		return
	}

	fmt.Println(string(personJSON))
	fmt.Println(string(personJSONB))

	// Create a new animal
	animalA := internal.Animal{Name: "Mateo", Age: 22, Race: "Galgo ingles"}

	// Print the name of the person and the animal
	animalA.PrintName()
	personA.PrintName()
}
