package engine

import (
	"eight-stones/ecs-tank-engine/engine/components"
)

// gameInfo inner info.
type gameInfo struct {
	NumberGamers        int
	MaxNumberGamers     int
	SizeX               int
	SizeY               int
	PreSelectPlaces     [][]int
	PreSelectDirections []components.Direction
}
