package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculoSalario_CategoriaA(t *testing.T) {
	// arrange

	minutosTrabajados := 60
	categoria := "A"

	expected := 4500.0
	// act
	result := calculoSalario(minutosTrabajados, categoria)

	// assert
	require.Equal(t, expected, result, "El salario es el esperado")
}

func TestCalculoSalario_CategoriaB(t *testing.T) {
	// arrange

	minutosTrabajados := 60
	categoria := "B"

	expected := 1800.0
	// act
	result := calculoSalario(minutosTrabajados, categoria)

	// assert
	require.Equal(t, expected, result, "El salario es el esperado")
}

func TestCalculoSalario_CategoriaC(t *testing.T) {
	// arrange

	minutosTrabajados := 60
	categoria := "C"

	expected := 1000.0
	// act
	result := calculoSalario(minutosTrabajados, categoria)

	// assert
	require.Equal(t, expected, result, "El salario es el esperado")
}
