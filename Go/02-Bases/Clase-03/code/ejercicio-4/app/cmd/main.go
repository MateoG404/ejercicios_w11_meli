package main

import (
	// Import the missing package
	"fmt"
)

func main() {
	fmt.Println("Hola Mundo")
	var operacion string
	var cantidadNumeros int

	fmt.Print("¿Qué operación desea hacer? ")
	fmt.Scanf("%s", &operacion)

	fmt.Print("¿Cuántos números desea ingresar? ")
	fmt.Scanf("%d", &cantidadNumeros)

	nums := make([]int, cantidadNumeros)
	for i := 0; i < cantidadNumeros; i++ {
		var num int
		fmt.Scanf("%d", &num)
		nums[i] = num
	}
	internal.Orchestator(operacion, nums...)
}
