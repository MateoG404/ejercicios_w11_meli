package internal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWhiteSharl(t *testing.T) {
	// Negative Test : shark is tired
	t.Run("shark is tired", func(t *testing.T) {
		// Arrange

		// - Create a shark that is tired
		shark := NewWhiteShark(true, true, 0)

		// - Create a tuna that is fast
		tuna := NewTuna("tuna", 1)

		// Act
		result := shark.Hunt(tuna)

		// Assert
		require.Equal(t, ErrSharkIsTired, result)
	})

	// Negative Test : shark already eaten a tuna
	t.Run("shark hunts", func(t *testing.T) {
		// Arrange
		// - Create a shark that is hungry and no
		shark := NewWhiteShark(false, false, 1)
		// - Create a tuna that is slow
		tuna := NewTuna("tuna", 1)

		// Act
		result := shark.Hunt(tuna)
		// Assert
		require.Equal(t, ErrSharkIsNotHungry, result)
	})

	// Negative Test : the tuna is faster than the shark

	t.Run("the tuna is faster than the shark", func(t *testing.T) {
		// Arrange

		// - Create a shark that is hungry and not tired
		shark := NewWhiteShark(true, false, 1)
		// - Create a tuna that is fast
		tuna := NewTuna("tuna", 2)
		// Act
		result := shark.Hunt(tuna)

		// Assert
		require.Equal(t, ErrSharkIsSlower, result)

	})

	// Positive Test : the shark hunts the tuna and the shark is not hungry anymore and is tired
	t.Run("the shark hunts the tuna and the shark is not hungry anymore and is tired", func(t *testing.T) {
		// Arrange

		// - Create a shark that is hungry and not tired
		shark := NewWhiteShark(true, false, 2)
		// - Create a tuna that is slow
		tuna := NewTuna("tuna", 1)
		// Act
		result := shark.Hunt(tuna)
		// Assert
		require.Nil(t, result) // no error

		require.False(t, shark.Hungry) // shark is not hungry anymore
		require.True(t, shark.Tired)   // shark is tired
	})

	// Unit Test : Tuna is nil
	t.Run("tuna is nil", func(t *testing.T) {

		// Arrange

		// - Create a shark that is hungry and not tired
		shark := NewWhiteShark(true, false, 2)
		// - Create a tuna that is nil
		var tuna *Tuna = nil
		// Act
		result := shark.Hunt(tuna)

		// Assert
		require.Equal(t, ErrTunaIsNil, result)
	})

}
