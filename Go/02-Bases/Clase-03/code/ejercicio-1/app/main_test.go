package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestImpuestoSalario_50000Low(t *testing.T) {
	// arrange
	salario := 10000.0
	expected := 10000.0

	// act
	result := impuestoSalario(salario)
	fmt.Println(result)
	// assert
	require.Equal(t, expected, result, "El salario es el esperado")
}
