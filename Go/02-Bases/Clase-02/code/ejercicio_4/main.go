package main

import "fmt"

func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

	fmt.Println("La edad de Benjamin es ", employees["Benjamin"])

	mayores := 0
	for i := range employees {
		if employees[i] > 21 {
			mayores++
		}
	}
	fmt.Println("La cantidad de empleados mayores de 21 es ", mayores)

	employees["Federico"] = 25
	delete(employees, "Pedro")

	fmt.Println("El nuevo map es", employees)

}
