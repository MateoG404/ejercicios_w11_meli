package main

import "fmt"

func main() {
	var palabra string

	fmt.Println("Ingrese una palabra para ser deletreada: ")
	fmt.Scanln(&palabra)

	fmt.Println("La cantidad de letras que tiene la palabra son: ", len(palabra))
	fmt.Println("La impresion de las letras de la palabra son :")

	for i := range palabra {
		fmt.Println(string(palabra[i]))
	}
}
