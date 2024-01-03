package main

import (
	"app_ejercicio4/app_ejercicio5/internal"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestAnimalDog test the function AnimalDog
func TestAnimalDog(t *testing.T) {

	// arrange

	animal := "dog"
	expected := "La cantidad de alimento que debe consumir es de 10kg"

	// act
	result := internal.Orchestator(animal)

	// assert

	require.Equal(t, expected, result())

}

// TestAnimalCat test the function AnimalCat

func TestAnimalCat(t *testing.T) {

	// arrange

	animal := "cat"
	expected := "La cantidad de alimento que debe consumir es de 5kg"

	// act
	result := internal.Orchestator(animal)

	// assert

	require.Equal(t, expected, result())

}

// TestAnimalHamster test the function AnimalHamster

func TestAnimalHamster(t *testing.T) {
	// arrange

	animal := "hamster"
	expected := "La cantidad de alimento que debe consumir es de 250gr"

	// act
	result := internal.Orchestator(animal)

	// assert

	require.Equal(t, expected, result())
}

// TestAnimalTarantula test the function AnimalTarantula
func TestAnimalTarantula(t *testing.T) {
	// arrange

	animal := "tarantula"
	expected := "La cantidad de alimento que debe consumir es de 150gr"

	// act
	result := internal.Orchestator(animal)

	// assert

	require.Equal(t, expected, result())
}
