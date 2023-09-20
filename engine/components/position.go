package components

import "eight-stones/ecs-tank-engine/engine/common"

// Position describes the location, i.e., the coordinates and direction of the entity.
type Position struct {
	X         int
	Y         int
	Direction common.Direction
}

// GetPosition returns Position component.
func (p *Position) GetPosition() *Position { return p }
