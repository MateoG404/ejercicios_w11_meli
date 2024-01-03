package tools_test

import (
	"intro-unit-testing/ej1/tools"
	"testing"
)

func TestElementAtIndex_ExistentIndex(t *testing.T) {

	// arrange : configurar el escenario
	slice := []int{1, 2, 3, 4, 5}
	index := 0

	expected := 1

	// act : ejecutar la acción
	result := tools.ElementAtIndex(slice, index)

	// assert : verificar el resultado
	// si el resultado es distinto al esperado, falla el test

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
		return
	}
	t.Log("success")
}

func TestElementAtIndex_NonExistingIndex(t *testing.T) {
	// arrange
	slice := []int{1, 2, 3, 4, 5}
	idx := len(slice) - 1

	expected := 5

	// act
	result := tools.ElementAtIndex(slice, idx)

	// assert
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
		return
	}
}
