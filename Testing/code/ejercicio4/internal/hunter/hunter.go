package hunter

import (
	"ejercicio4/internal/positioner"
	"ejercicio4/internal/prey"
	"errors"
)

var (
	// ErrCanNotHunt is returned when the hunter can not hunt the prey
	ErrCanNotHunt = errors.New("can not hunt the prey")
)

// Hunter is an interface that represents a hunter
type Hunter interface {
	// Hunt hunts the prey
	Hunt(prey prey.Prey) (duration float64, err error)
	// Configure configures the hunter
	Configure(speed float64, position *positioner.Position)
}
