package main

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	green = color.New(color.FgGreen).SprintFunc()
)

func main() {
	fmt.Println(green("Inicio del programa"))

	fmt.Println(green("Finalizacion del programa"))

}
