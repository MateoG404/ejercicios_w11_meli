package main

import (
	"fmt"
)

func main() {
	var mes int
	fmt.Print("Ingrese el numero del mes: ")
	fmt.Scan(&mes)

	switch mes {
	case 1:
		fmt.Println("El mes es Enero")
	case 2:
		fmt.Println("El mes es Febrero")
	case 3:
		fmt.Println("El mes es Marzo")
	case 4:
		fmt.Println("El mes es Abril")
	case 5:
		fmt.Println("El mes es Mayo")
	case 6:
		fmt.Println("El mes es Junio")
	case 7:
		fmt.Println("El mes es Julio")
	case 8:
		fmt.Println("El mes es Agosto")
	case 9:
		fmt.Println("El mes es Septiembre")
	case 10:
		fmt.Println("El mes es Octubre")
	case 11:
		fmt.Println("El mes es Noviembre")
	case 12:
		fmt.Println("El mes es Diciembre")

	}

}
