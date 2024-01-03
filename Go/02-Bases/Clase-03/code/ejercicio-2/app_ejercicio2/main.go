package main

import (
	"fmt"
	"strconv"
)

func promedio(grades ...float64) (mean float64) {
	sum := 0.0
	for _, grade := range grades {
		sum += grade
	}
	mean = sum / float64(len(grades))
	return mean
}

func main() {

	// Ingresar la cantidad de alumnos
	fmt.Println("¿Cuantos alumnos desea ingresar?")
	var i int
	// Leer la cantidad de alumnos
	fmt.Scanln(&i)

	// Ingresar las notas de cada alumno
	for j := 0; j < i; j++ {
		fmt.Println("Para ingresar las notas de un nuevo alumno ingrese la letra S, para salir ingrese la letra N")

		var temp string
		grades := []float64{}
		// Ingresar las notas de un alumno hasta que el usuario diga que no mas
		for {
			// Leer la nota
			fmt.Scanln(&temp)
			// Si el usuario ingresa la letra N, salir del ciclo para leer nuevo alumno
			if temp == "N" {
				break
			}
			// Si el usuario ingresa la letra S, continuar con el ciclo para ingresar notas

			grade, err := strconv.ParseFloat(temp, 64)
			if err != nil {
				fmt.Println("Error: Ingrese un número válido.")
				for {
					fmt.Scanln(&temp)
					grade, err := strconv.ParseFloat(temp, 64)
					if err != nil {
						fmt.Println("Error: Ingrese un número válido.")
					} else {
						grades = append(grades, grade)
						break
					}
				}
			} else {
				grades = append(grades, grade)
			}
		}
		fmt.Println("El promedio para el alumno es :", promedio(grades...))

	}
}
