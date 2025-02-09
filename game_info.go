package engine

import (
	"github.com/Eight-Stones/ecs-tank-engine/components"
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
