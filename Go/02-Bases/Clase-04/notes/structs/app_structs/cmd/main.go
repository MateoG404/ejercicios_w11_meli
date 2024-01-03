package main

import (
	"app_structs/app_structs/internal"
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Using structs")

	// Create a new person
	personA := internal.Person{FirstNmame: "Mateo", LastName: "Gutierrez", CellPhone: "123456789", Address: "Street 123"}
	fmt.Println(personA)

	// Store the person in a JSON file
	personJSON, err := json.Marshal(personA)
	if err != nil {
		fmt.Println("Error marshaling person:", err)
		return
	}
	fmt.Println(string(personJSON))
}
