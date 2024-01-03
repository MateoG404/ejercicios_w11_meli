package main

import "fmt"

func impuestoSalario(salario float64) float64 {
	if salario > 50000 {
		return salario - (salario * 0.17)
	} else if salario > 150000 {
		return salario - (salario * 0.27)
	}
	return salario
}

func main() {
	var salario float64
	fmt.Println("Ingrese el salario de su empleado:")
	fmt.Scan(&salario)

	fmt.Println("El salario del empleado es:", impuestoSalario(salario))
}
