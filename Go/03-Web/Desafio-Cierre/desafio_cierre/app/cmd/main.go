package main

import (
	"desafio_cierre/app/platform"
	"fmt"

	"github.com/fatih/color"
)

var (
	green = color.New(color.FgGreen).SprintFunc()
	red   = color.New(color.FgRed).SprintFunc()
)

func main() {
	// Start the application
	fmt.Println(green("Inicio de la aplicación"))

	// Configurate the environment variables
	// Create and configurate the server
	app := platform.NewServerChi(":8080")

	// Start the server
	if err := app.Start(); err != nil {
		fmt.Println(red("Error al iniciar el servidor: %s\n", err.Error()))
	}

}
