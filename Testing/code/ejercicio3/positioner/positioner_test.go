// This script contains the unit tests for the positioner package
package positioner

import (
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
