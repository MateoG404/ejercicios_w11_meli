package hunter_test

import (
	"ejercicio3/hunter"
	"ejercicio3/positioner"
	"ejercicio3/prey"
	"ejercicio3/simulator"

	"testing"

	"github.com/stretchr/testify/require"
)

// Tests for the hunter.NewWhiteShark(implementation of the Hunter interface
func TestWhiteSharkHunt(t *testing.T) {
	t.Run("white shark hunts a prey - has speed and short distance", func(t *testing.T) {
		// arrange
		pr := prey.NewPreyStub()
		pr.GetPositionFunc = func() (position *positioner.Position) {
			return &positioner.Position{X: 0, Y: 0, Z: 0}
		}
		pr.GetSpeedFunc = func() (speed float64) {
			return 10
		}

		sm := simulator.NewCatchSimulatorMock()
		sm.CanCatchFunc = func(hunter, prey *simulator.Subject) (canCatch bool) {
			return true
		}

		impl := hunter.NewWhiteShark(100, &positioner.Position{X: 1, Y: 1, Z: 1}, sm)

		// act
		err := impl.Hunt(pr)

		// assert
		expErr := error(nil)
		expMockCallCanCatch := 1
		require.ErrorIs(t, err, expErr)
		require.Equal(t, expMockCallCanCatch, sm.Calls.CanCatch)
	})

	t.Run("white shark can not hunt a prey - has short speed", func(t *testing.T) {
		// arrange
		pr := prey.NewPreyStub()
		pr.GetPositionFunc = func() (position *positioner.Position) {
			return &positioner.Position{X: 0, Y: 0, Z: 0}
		}
		pr.GetSpeedFunc = func() (speed float64) {
			return 10
		}

		sm := simulator.NewCatchSimulatorMock()
		sm.CanCatchFunc = func(hunter, prey *simulator.Subject) (canCatch bool) {
			return false
		}

		impl := hunter.NewWhiteShark(1, &positioner.Position{X: 1, Y: 1, Z: 1}, sm)

		// act
		err := impl.Hunt(pr)

		// assert
		expErr := hunter.ErrCanNotHunt
		expErrMsg := "can not hunt the prey: shark can not catch the prey"
		expMockCallCanCatch := 1
		require.ErrorIs(t, err, expErr)
		require.EqualError(t, err, expErrMsg)
		require.Equal(t, expMockCallCanCatch, sm.Calls.CanCatch)
	})

	t.Run("white shark can not hunt a prey - has long distance", func(t *testing.T) {
		// arrange
		pr := prey.NewPreyStub()
		pr.GetPositionFunc = func() (position *positioner.Position) {
			return &positioner.Position{X: 0, Y: 0, Z: 0}
		}
		pr.GetSpeedFunc = func() (speed float64) {
			return 10
		}

		sm := simulator.NewCatchSimulatorMock()
		sm.CanCatchFunc = func(hunter, prey *simulator.Subject) (canCatch bool) {
			return false
		}

		impl := hunter.NewWhiteShark(100, &positioner.Position{X: 1000, Y: 1000, Z: 1000}, sm)

		// act
		err := impl.Hunt(pr)

		// assert
		expErr := hunter.ErrCanNotHunt
		expErrMsg := "can not hunt the prey: shark can not catch the prey"
		expMockCallCanCatch := 1
		require.ErrorIs(t, err, expErr)
		require.EqualError(t, err, expErrMsg)
		require.Equal(t, expMockCallCanCatch, sm.Calls.CanCatch)
	})
}

// Test for CreateWhiteShark
func TestCreateWhiteShark(t *testing.T) {

	// Case 1 : Sucessful creation of a WhiteShark with default speed
	t.Run("Case 1 : Sucessful creation of a WhiteShark with default speed", func(t *testing.T) {
		// Arrange
		// - Create a CatchSimulatorMock
		simulator := simulator.NewCatchSimulatorMock()

		// Act
		// .- Create a WhiteShark with Default speed
		hunter := hunter.CreateWhiteShark(simulator)

		// Assert
		// -Get the speed of the WhiteShark
		speed := hunter.GetSpeed()
		require.True(t, 15.0 <= speed && speed <= 159.0, "speed is not between 15 and 159")
	})

	// Case 2 : Sucessful creation of a WhiteShark with default position
	t.Run("Case 2 : Sucessful creation of a WhiteShark with default position", func(t *testing.T) {
		// Arrange
		// - Create a CatchSimulatorMock
		simulator := simulator.NewCatchSimulatorMock()

		// Act
		// - Create a WhiteShark with Default position
		hunter := hunter.CreateWhiteShark(simulator)

		// Assert
		// - Get the position of the WhiteShark
		position := hunter.GetPosition()
		require.True(t, 0.0 <= position.X && position.X <= 500 && 0.0 <= position.Y && position.Y <= 500 && 0.0 <= position.Z && position.Z <= 500, "position is not between 0 and 500")
	})

	// Case 3 : Sucessful creation of a WhiteShark with default speed and position
	t.Run("Case 3 : Sucessful creation of a WhiteShark with default speed and position", func(t *testing.T) {
		// Arrange
		// - Create a CatchSimulatorMock
		simulator := simulator.NewCatchSimulatorMock()

		// Act
		// - Create a WhiteShark with Default speed and position
		hunter := hunter.CreateWhiteShark(simulator)

		// Assert
		// - Get the speed of the WhiteShark
		speed := hunter.GetSpeed()
		// - Get the position of the WhiteShark
		position := hunter.GetPosition()

		// Assert
		// - Check that the speed is between 15 and 159
		require.True(t, 15.0 <= speed && speed <= 159.0, "speed is not between 15 and 159")
		// - Check that the position is between 0 and 500
		require.True(t, 0.0 <= position.X && position.X <= 500 && 0.0 <= position.Y && position.Y <= 500 && 0.0 <= position.Z && position.Z <= 500, "position is not between 0 and 500")
	})
}

// Test for HunterStubs

func TestHunterStub(t *testing.T) {
	// Case 1 : Sucessful creation of a HunterStub
	t.Run("Case 1 : Sucessful creation of a HunterStub ", func(t *testing.T) {
		// Arrange
		// ...

		// Act
		// - Create a HunterStub
		hunter := hunter.NewHunterStub()
		// Assert
		// - Get the speed of the HunterStub
		require.NotNil(t, hunter, "HunterStub is not created")
		require.Nil(t, hunter.GetSpeedFunc, "GetSpeedFunc is not nil")
		require.Nil(t, hunter.GetPositionFunc, "GetPositionFunc is not nil")
	})
}
