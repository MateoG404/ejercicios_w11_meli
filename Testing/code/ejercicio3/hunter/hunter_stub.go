// This script is a stub for the hunter package
package hunter

import "ejercicio3/positioner"

// NewHunterStub creates a new HunterStub
func NewHunterStub() (hunter *HunterStub) {
	hunter = &HunterStub{}
	return
}

// Struct HunterStub is a stub for Hunter
type HunterStub struct {
	// GetSpeedFunc externalize the GetSpeed method
	GetSpeedFunc func() (speed float64)
	// GetPositionFunc externalize the GetPosition method
	GetPositionFunc func() (position *positioner.Position)
}

// GetSpeed
func (s *HunterStub) GetSpeed() (speed float64) {
	return s.GetSpeedFunc()
}

// GetPosition
func (s *HunterStub) GetPosition() (position *positioner.Position) {
	return s.GetPositionFunc()
}
