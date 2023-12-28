package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("Hello, world!")
	fmt.Println("The current time is", time.Now())

	var numMateo = 10
	var floatMateo = float32(numMateo)

	println(floatMateo)

	fmt.Println(Sum(1, 2))
	color.Red("Hello, world!")

}

func Sum(a, b int) int {
	return a + b
}
