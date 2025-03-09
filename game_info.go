package engine

import (
	"github.com/Eight-Stones/ecs-tank-engine/v2/components"
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
