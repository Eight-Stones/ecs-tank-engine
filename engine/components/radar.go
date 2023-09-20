package components

// Radar adds big view.
type Radar struct {
	Radius int
	*Recharge
}

// GetRadar returns Radar component.
func (r *Radar) GetRadar() *Radar { return r }
