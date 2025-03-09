package systems

import (
	"github.com/Eight-Stones/ecs-tank-engine/v2/components"
)

var directionIncs = map[components.Direction][]int{
	components.Left:  {-1, 0},
	components.Right: {1, 0},
	components.Down:  {0, -1},
	components.Up:    {0, 1},
}

// PositionSystem describes position system.
type PositionSystem interface {
	GetPosition() *components.Position
}

// GetIncrementPosition return increment for coordinate by position.
func GetIncrementPosition(direction components.Direction) (x, y int) {
	values := directionIncs[direction]
	return values[0], values[1]
}

// ChangePosition changes position on select x/y coordinate.
func ChangePosition(incX, incY int, in PositionSystem) {
	position := in.GetPosition()
	position.X += incX
	position.Y += incY
}

// IsCollision checks that first object collide with second.
func IsCollision(first PositionSystem, second PositionSystem) int {
	if first == second {
		return Fail
	}

	incX, incY := GetIncrementPosition(first.GetPosition().Direction)
	pMove := first.GetPosition()
	pCllsn := second.GetPosition()

	if pCllsn.X == pMove.X+incX && pCllsn.Y == pMove.Y+incY {
		return Success
	}

	return Fail
}
