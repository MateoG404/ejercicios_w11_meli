package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPromedio(t *testing.T) {
	// arrange
	grades := []float64{10, 10, 10, 10, 10}
	expected := 10.0

	// act
	result := promedio(grades...)
	// assert

	require.Equal(t, expected, result, "El promedio es el esperado")
}
