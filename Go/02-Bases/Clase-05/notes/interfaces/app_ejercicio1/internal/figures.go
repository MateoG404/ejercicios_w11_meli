// Code created for implement the interfaces in Go to calculate the area and perimeter for two figures

package internal

import (
	"math"
	"strconv"
)

// Figure is an interface to calculate the area and perimeter for two figures
type Figure interface {
	CalculateArea() float64
	CalculatePerimeter() float64
	InformationFigure() string
}

// Rectangle is a struct to represent a rectangle

type Rectangle struct {
	Width     float64
	Height    float64
	Area      float64
	Perimeter float64
}

// Area is a method to calculate the area for a rectangle
func (r *Rectangle) CalculateArea() {
	r.Area = r.Width * r.Height
}

// Perimeter is a method to calculate the perimeter for a rectangle
func (r *Rectangle) CalculatePerimeter() {
	r.Perimeter = 2 * (r.Width + r.Height)
}

// Information is a method to return the information for a rectangle
func (r *Rectangle) InformationFigure() string {
	return "Hola soy un Rectangulo con un ancho de " + strconv.FormatFloat(r.Width, 'f', -1, 64) + " y un alto de " + strconv.FormatFloat(r.Height, 'f', -1, 64) + " y un area de " + strconv.FormatFloat(r.Area, 'f', -1, 64) + " y un perimetro de " + strconv.FormatFloat(r.Perimeter, 'f', -1, 64)

}

// Circle is a struct to represent a circle

type Circle struct {
	Radius    float64
	Area      float64
	Perimeter float64
}

// Area is a method to calculate the area for a circle
func (c *Circle) CalculateArea() {
	c.Area = math.Pow(c.Radius, 2) * 3.14

}

// Perimeter is a method to calculate the perimeter for a circle
func (c *Circle) CalculatePerimeter() {
	c.Perimeter = 2 * c.Radius * 3.14

}

// Information is a method to return the information for a circle

func (c *Circle) InformationFigure() string {

	return "Hola soy un Circulo con un radio de " + strconv.FormatFloat(c.Radius, 'f', -1, 64) + " y un area de " + strconv.FormatFloat(c.Area, 'f', -1, 64) + " y un perimetro de " + strconv.FormatFloat(c.Perimeter, 'f', -1, 64)

}
