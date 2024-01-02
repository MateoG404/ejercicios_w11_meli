package internal

import (
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func Orchestator(operation string, nums ...int) {
	// Implementacion del orquestador
	var result int
	switch operation {
	case minimum:
		result = Minimum(nums...)
		fmt.Println(result)
	case average:
		result = Average(nums...)
		fmt.Println(result)
	case maximum:
		result = Maximum(nums...)
		fmt.Println(result)
	default:
		fmt.Println("No se encontro la operacion")
	}
}
