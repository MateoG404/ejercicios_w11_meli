package main

import "fmt"

func main() {
	// A continuación se muestran las correciones de las declaraciones de variables:

	// Se usaran declaraciones implicitas para las variables

	lastName := "Smith"

	age := 35

	boolean := "false"

	salary := 45857.90

	firstName := "Mary"

	fmt.Println(lastName, age, boolean, salary, firstName)

	if sum := age + int(salary); sum > 100 {
		fmt.Println("La suma es mayor a 100")
	}
}
