package components

// Vision adds view.
type Vision struct {
	Radius int
	*Recharge
}

// GetVision returns Vision component.
func (v *Vision) GetVision() *Vision { return v }
