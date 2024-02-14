// This script contains the unit tests for the positioner package
package positioner

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

// Unit tests for the Position struct
func TestGetLinearDistance(t *testing.T) {

	// Case 1: The coordinates are negative
	t.Run("Case 1: the coordinates are negative", func(t *testing.T) {
		// Arrange
		// - Set the positioner
		positioner := NewPositionerDefault()
		// - Set the from position
		from := &Position{
			X: -1,
			Y: -1,
			Z: -1,
		}
		// - Set the to position
		to := &Position{
			X: -1,
			Y: -1,
			Z: -1,
		}

		// Act
		// - Get the linear distance
		linearDistance := positioner.GetLinearDistance(from, to)

		// Assert
		// - Check the linear distance
		require.Equal(t, 0.0, linearDistance)

	})

	// Case 2: The coordinates are positive
	t.Run("Case 2: the coordinates are positive", func(t *testing.T) {
		// Arrange
		// - Set the positioner
		positioner := NewPositionerDefault()
		// - Set the from position
		from := &Position{
			X: 1,
			Y: 1,
			Z: 1,
		}
		// - Set the to position
		to := &Position{
			X: 1,
			Y: 1,
			Z: 1,
		}
		// Act
		// - Get the linear distance
		linearDistance := positioner.GetLinearDistance(from, to)

		// Assert
		// - Check the linear distance
		require.Equal(t, 0.0, linearDistance)

	})

	// Case 3: The coordinates returns a linear distance whitout decimals
	t.Run("Case 3: the coordinates returns a linear distance whitout decimals", func(t *testing.T) {
		// Arrange
		// - Set the positioner
		positioner := NewPositionerDefault()
		// - Set the from position
		from := &Position{
			X: 1,
			Y: 1,
			Z: 12}
		// - Set the to position
		to := &Position{
			X: 1,
			Y: 1,
			Z: 9,
		}
		// Act
		// - Get the linear distance
		linearDistance := positioner.GetLinearDistance(from, to)

		// Assert
		// - Check the linear distance
		require.Equal(t, 3.0, linearDistance)
	})
}

// Test the NewPositionerStub method
func TestNewPositionerStub(t *testing.T) {
	// Case 1 : Sucessful creation of a PositionerStub
	t.Run("Case 1 : Sucessful creation of a PositionerStub", func(t *testing.T) {
		// Arrange
		// ...

		// Act
		// Creation a new PositionerStub
		positioner := NewPositionerStub()

		// Assert
		// Check the positioner is not nil
		require.NotNil(t, positioner, "The positioner is not nil")

	})
}

// Test GetLinearDistanceMock test the simualtor mock

func TestGetLinearPositionerStub(t *testing.T) {
	// Case 1 : Sucessful creation of GetLinearDistanceMock
	t.Run("Case 1 : Sucessful creation of GetLinearDistanceMock", func(t *testing.T) {
		// Arrange
		// - Creation PositionerStub
		positioner := NewPositionerStub()

		// - Set the from and to values
		from := &Position{
			X: 1,
			Y: 1,
			Z: 1,
		}
		to := &Position{
			X: 0,
			Y: 0,
			Z: 0}

		// Act
		// - Calls the GetLinearDistanceMock method
		positioner.GetLinearDistanceFunc = func(from, to *Position) float64 {
			return math.Sqrt(3)
		}
		linearDistance := positioner.GetLinearDistance(from, to)
		// Assert
		expectedDistance := math.Sqrt(3) // La distancia esperada es la raíz cuadrada de 3
		require.Equal(t, expectedDistance, linearDistance)

	})
}
