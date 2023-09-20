package engine

import "eight-stones/ecs-tank-engine/engine/common"

// GameInfo inner info.
type GameInfo struct {
	NumberGamers        int
	MaxNumberGamers     int
	SizeX               int
	SizeY               int
	PreSelectPlaces     [][]int
	PreSelectDirections []common.Direction
}
