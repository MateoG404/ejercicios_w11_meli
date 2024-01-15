package main

import (
	"fmt"
	application "patch_put_delete/app/internal/application"

	"github.com/fatih/color"
)

var (
	green = color.New(color.FgGreen).SprintFunc()
)

func main() {
	// Configurate the environment variables

	// Run application
	fmt.Println(green("Application runnings"))
	// env
	// ...

	// app
	// - config
	app := application.NewServerChi("8080")
	// - run
	if err := app.Run(); err != nil {
		fmt.Println("Error running the application")
		fmt.Println(err)
		return
	}
	fmt.Println("aca")
}
