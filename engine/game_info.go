package engine

import "eight-stones/ecs-tank-engine/engine/common"

// gameInfo inner info.
type gameInfo struct {
	NumberGamers        int
	MaxNumberGamers     int
	SizeX               int
	SizeY               int
	PreSelectPlaces     [][]int
	PreSelectDirections []common.Direction
}
