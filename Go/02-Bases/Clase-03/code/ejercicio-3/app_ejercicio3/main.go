package main

import "fmt"

func calculoSalario(minutesMonth int, category string) (salary float64) {
	hours := float64(minutesMonth) / 60
	switch category {
	case "C":
		salary = float64(hours * 1000)
	case "B":
		salary = float64((hours * 1500))
		salary += (salary * 0.2)
	case "A":
		fmt.Println("hours", hours)
		salary = float64((hours * 3000))
		salary += (salary * 0.5)
	default:
		salary = 0
	}
	return salary
}

func main() {
	// Leer la cantidad de alumnos

	// Ingresar la cantidad de empleados
	fmt.Println("¿Cuantos empleados desea ingresar?")

	var empleados int
	var err error
	// Leer la cantidad de empleados
	fmt.Scanln(&empleados)

	// Ingresar las notas de cada empleado
	for i := 0; i < empleados; i++ {

		var minutesMonth int
		var category string
		var salary float64

		fmt.Println("Ingrese los minutos trabajos por mes")
		// Ingresar los minutos trabajados por mes
		_, err = fmt.Scanln(&minutesMonth)

		// Validar que el usuario ingrese un número para los minutos trabajados por mes
		if err != nil {
			for {
				fmt.Println("Error: Ingrese un número válido.")
				_, err = fmt.Scanln(&minutesMonth)
				// Si el usuario ingresa un número, salir del ciclo
				if err == nil {
					break
				}
			}
		}

		// Ingresar la categoria del empleado
		fmt.Println("Ingrese la categoria del empleado")
		_, _ = fmt.Scanln(&category)

		salary = calculoSalario(minutesMonth, category)

		fmt.Println("El salario para el empleado es :", salary)
	}
}
