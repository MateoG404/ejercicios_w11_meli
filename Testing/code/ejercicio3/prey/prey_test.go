package prey

import (
	"ejercicio3/positioner"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestGetSpeed tests the GetSpeed method

func TestGetSpeed(t *testing.T) {
	// Case 1 : The speed is default
	t.Run("Case 1 : The speed is default", func(t *testing.T) {
		// Arrange
		// - Create and set the tuna
		tuna := CreateTuna()

		// Act
		// - Get the speed
		speed := tuna.GetSpeed()

		// Assert
		// - Check the speed is default between 15 and 252
		require.True(t, speed >= 15.0 && speed <= 252.0, "The speed is not default")
	})

	// Case 2 : The speed is not default
	t.Run("Case 2 : The speed is not default", func(t *testing.T) {
		// Arrange

		// - Create and set the tuna
		tuna := NewTuna(100.0, nil)

		// Act
		// - Get the speed
		speed := tuna.GetSpeed()

		// Assert
		// - Check the speed is not default
		require.Equal(t, 100.0, speed, "The speed is not default")

	})
}

// TestGetPosition tests the GetPosition method

func TestGetPosition(t *testing.T) {
	// Case 1 : The position is default
	t.Run("Case 1 : The position is default", func(t *testing.T) {
		// Arrange
		// - Create and set the tuna with default position
		tuna := CreateTuna()

		// Act
		// - Get the position
		position := tuna.GetPosition()

		// Assert
		require.True(t, position.X >= 0.0 && position.X <= 500.0, "The position is not default")
	})

	// Case 2 : The position is not default
	t.Run("Case 2 : The position is not default", func(t *testing.T) {
		// Arrange
		// - Create and set the tuna with not default position
		tuna := NewTuna(0.0, &positioner.Position{
			X: 100.0,
			Y: 100.0,
			Z: 100.0})

		// Act
		// - Get the position
		position := tuna.GetPosition()

		// Assert
		// - Check the position is not default
		require.Equal(t, 100.0, position.X, "The position is not default")
	})
}
