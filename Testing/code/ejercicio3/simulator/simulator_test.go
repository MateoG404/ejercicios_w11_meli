package simulator

import (
	"ejercicio3/positioner"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestCatchSimulator tests the CatchSimulator

func TestCatchSimulator(t *testing.T) {
	// Case 1 : The hunter can catch the prey

	t.Run("Case 1 : The hunter can catch the prey", func(t *testing.T) {
		// Arrange
		// - Create a hunter and set the position and speed
		hunter := &Subject{
			Position: &positioner.Position{
				X: 0,
				Y: 0,
				Z: 0,
			}, Speed: 10,
		}

		// - Create a prey and set the position and speed

		prey := &Subject{
			Position: &positioner.Position{
				X: 1,
				Y: 1,
				Z: 1,
			}, Speed: 1}

		// Act
		// - Create a catch simulator and set the max time to catch
		simulator := NewCatchSimulatorDefault(1, positioner.NewPositionerDefault())
		// - Call the CanCatch method
		canCatch := simulator.CanCatch(hunter, prey)

		// Assert
		// - Check if the hunter can catch the prey
		require.True(t, canCatch)
	})

	// Case 2 : The hunter can't catch the prey because its slow
	t.Run("Case 2 : The hunter can't catch the prey because its slow", func(t *testing.T) {
		// Arrange
		// - Create a hunter and set the position and speed
		hunter := &Subject{
			Position: &positioner.Position{
				X: 0,
				Y: 0,
				Z: 0,
			}, Speed: 1}

		// - Create a prey and set the position and speed
		prey := &Subject{
			Position: &positioner.Position{
				X: 2,
				Y: 2,
				Z: 2}, Speed: 10}

		// - Create a catch simulator and set the max time to catch
		simulator := NewCatchSimulatorDefault(1, positioner.NewPositionerDefault())

		// Act
		// Call the CanCatch method
		canCatch := simulator.CanCatch(hunter, prey)

		// Assert
		// Check if the hunter can't catch the prey
		require.False(t, canCatch)
	})
	// Case 3 : The hunter can't catch the prey because its faar but its fast
	t.Run("Case 3 : The hunter can't catch the prey because its faar but its fast", func(t *testing.T) {
		// Arrange
		// - Create a hunter and set the position and speed
		hunter := &Subject{
			Position: &positioner.Position{
				X: 0,
				Y: 0,
				Z: 0,
			}, Speed: 10}
		// - Create a prey and set the position and speed
		prey := &Subject{
			Position: &positioner.Position{
				X: 100,
				Y: 100,
				Z: 100,
			}, Speed: 1}

		// - Create a catch simulator and set the max time to catch
		simulator := NewCatchSimulatorDefault(1, positioner.NewPositionerDefault())

		// Act
		// - Call the CanCatch method
		canCatch := simulator.CanCatch(hunter, prey)

		// Assert
		// - Check if the hunter can't catch the prey
		require.False(t, canCatch)
	})
}
