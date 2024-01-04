// This code is created to implement the interfaces notes about the interfaces in Go
package main

import (
	"ejercicio_1/app_ejercicio1/internal"
	"fmt"
)

func main() {

	// Create Circle
	circle := internal.Circle{Radius: 10}

	// Create Rectangle
	rectangle := internal.Rectangle{Width: 10, Height: 5}

	//  Calculate the area and perimeter for those figures

	circle.CalculateArea()
	rectangle.CalculateArea()

	// Calculate the perimeter for those figures
	circle.CalculatePerimeter()
	rectangle.CalculatePerimeter()

	// Print the information for those figures

	fmt.Println("La informacion de las figuras es :")
	fmt.Println(circle.InformationFigure())
	fmt.Println(rectangle.InformationFigure())

}
