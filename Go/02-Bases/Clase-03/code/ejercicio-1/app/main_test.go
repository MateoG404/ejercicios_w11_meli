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

func TestImpuestoSalario_50000High(t *testing.T) {

	// arrange
	salario := 50001.0
	expected := 41500.83

	// act
	result := impuestoSalario(salario)
	// assert
	require.Equal(t, expected, result, "El salario es el esperado")
}

func TestImpuestoSalario_150000High(t *testing.T) {
	// arrange
	salario := 150001.0
	expected := 109500.73

	// act
	result := impuestoSalario(salario)

	// assert
	require.Equal(t, expected, result, "El salario es el esperado")

}
