package main

import (
	"fmt"
)

func main() {

	// Creacion de un map

	myMap := map[string]string{"banano": "amarillo", "manzana": "rojo", "pera": "verde"}
	fmt.Println(myMap)

	// Creacion de un map vacio con make

	myMap2 := make(map[string]string)
	fmt.Println(myMap2)

	for key, element := range myMap {
		fmt.Println("La clave es", key, "y el elemento es", element)
	}
}
