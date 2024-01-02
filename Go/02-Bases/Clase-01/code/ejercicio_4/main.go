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

	// Ejemplo para crear un slice de enteros
	//enteros := []int{1, 2, 3, 4, 4, 5, 6, 6, 6, 8, 9, 10}
	matr := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	for i, item := range matr {
		fmt.Println(i, item)
	}
	a := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, item := range a {
		fmt.Println(i, item)
	}
}
