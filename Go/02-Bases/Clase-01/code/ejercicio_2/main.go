package main

import "fmt"

func main() {

	var temperatura, humedad, presion float64

	fmt.Println("Ingrese la temperatura, humedad y presion en ese mismo orden. ")
	fmt.Scan(&temperatura, &humedad, &presion)

	fmt.Println("La temperatura es", temperatura, "grados, la humedad es", humedad, "y la presion es", presion)

}
