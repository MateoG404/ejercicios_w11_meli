package positioner

// NewPositionerStub creates a new PositionerStub
func NewPositionerStub() (positioner *PositionerStub) {
	positioner = &PositionerStub{}
	return
}

// PositionerStub is a stub for Positioner
type PositionerStub struct {
	// GetLinearDistanceFunc externalize the GetLinearDistance method
	GetLinearDistanceFunc    func(from, to *Position) (linearDistance float64)
	NewPositionerDefaultFunc func() (positioner *Positioner)
}

// GetLinearDistance
func (s *PositionerStub) GetLinearDistance(from, to *Position) (linearDistance float64) {
	linearDistance = s.GetLinearDistanceFunc(from, to)
	return
}

// NewPositionerDefault
func (s *PositionerStub) NewPositionerDefault() (positioner *Positioner) {
	positioner = s.NewPositionerDefaultFunc()
	return
}
