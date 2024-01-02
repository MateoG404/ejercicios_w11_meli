package main

import "fmt"

func main() {

	var (
		edad         int
		statusEmpleo bool
		antiguedad   int
		sueldo       int64
	)

	fmt.Println("Por favor ingrese los siguientes campos:")

	fmt.Print("Edad: ")
	fmt.Scanf("%d\n", &edad)

	fmt.Print("Status Empleo: ")
	fmt.Scanf("%t\n", &statusEmpleo)

	fmt.Print("Antiguedad: ")
	fmt.Scanf("%d\n", &antiguedad)

	fmt.Print("Sueldo: ")
	fmt.Scanf("%d\n", &sueldo)

	switch {
	case edad > 22 && statusEmpleo == true && antiguedad > 1:
		fmt.Println("Usted tiene los requsitos para el prestamo")
		if sueldo > 1000000 {
			fmt.Println("Su prestamo no tendra cobro de intereses")
		}

	default:
		fmt.Println("Usted no cumple con los requisitos para el prestamo")
	}
}
