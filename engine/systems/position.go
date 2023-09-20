package systems

import (
	"eight-stones/ecs-tank-engine/engine/common"
	"eight-stones/ecs-tank-engine/engine/components"
)

// PositionSystem describes position system.
type PositionSystem interface {
	GetPosition() *components.Position
}

// GetIncrementPosition return increment for coordinate by position.
func GetIncrementPosition(direction common.Direction) (x, y int) {
	values := common.MovementValue[direction]
	return values[0], values[1]
}

// ChangePosition changes position on select x/y coordinate.
func ChangePosition(incX, incY int, in PositionSystem) {
	position := in.GetPosition()
	position.X += incX
	position.Y += incY
}

// CheckCollision checks that first object collide with second.
func CheckCollision(first PositionSystem, second PositionSystem) int {
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
