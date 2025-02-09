package components

// Position describes the location, i.e., the coordinates and direction of the entity.
type Position struct {
	X         int
	Y         int
	Direction Direction
}

// GetPosition returns Position component.
func (p *Position) GetPosition() *Position { return p }
