package hunter

import (
	"ejercicio3/positioner"
	"ejercicio3/prey"
	"errors"
)

var (
	// ErrCanNotHunt is returned when the hunter can not hunt the prey
	ErrCanNotHunt = errors.New("can not hunt the prey")
)

// Hunter is an interface that represents a hunter
type Hunter interface {
	// Hunt hunts the prey
	Hunt(prey prey.Prey) (err error)
	// GetSpeed returns the speed of the hunter in m/s
	GetSpeed() (speed float64)
	// GetPosition returns the position of the hunter
	GetPosition() (position *positioner.Position)
}
