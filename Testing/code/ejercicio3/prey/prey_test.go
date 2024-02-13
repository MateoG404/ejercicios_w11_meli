package prey

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
		assert.True(t, speed >= 15.0 && speed <= 252.0, "The speed is not default")
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
		assert.Equal(t, 100.0, speed, "The speed is not default")

	})
}

// TestGetPosition tests the GetPosition method
