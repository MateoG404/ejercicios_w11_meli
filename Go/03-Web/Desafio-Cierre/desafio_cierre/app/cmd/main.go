package main

import (
	"desafio_cierre/app/platform"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

var (
	green = color.New(color.FgGreen).SprintFunc()
	red   = color.New(color.FgRed).SprintFunc()
)

func main() {
	// Use the environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the environment variables
	addr := os.Getenv("ADDR")

	// Start the application
	fmt.Println(green("Inicio de la aplicación"))

	// Configurate the environment variables
	// Create and configurate the server
	app := platform.NewServerChi(addr)

	// Start the server
	if err := app.Start(); err != nil {
		fmt.Println(red("Error al iniciar el servidor: %s\n", err.Error()))
	}
}
