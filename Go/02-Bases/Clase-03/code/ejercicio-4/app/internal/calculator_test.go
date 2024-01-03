package internal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinimum(t *testing.T) {
	// arrange
	nums := []int{1, 2, 3, 4, 5}
	expected := 1
	// act
	result := Minimum(nums...)
	// assert
	require.Equal(t, expected, result, "El mínimo es el esperado")

}

func TestAverage(t *testing.T) {
	// arrange
	nums := []int{1, 2, 3, 4, 5}
	expected := 3
	// act
	result := Average(nums...)
	// assert
	require.Equal(t, expected, result, "El promedio es el esperado")
}

func TestMaximum(t *testing.T) {
	// arrange
	nums := []int{1, 2, 3, 4, 5}
	expected := 5
	// act
	result := Maximum(nums...)
	// assert
	require.Equal(t, expected, result, "El máximo es el esperado")

}
