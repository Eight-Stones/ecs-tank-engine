package systems

import "eight-stones/ecs-tank-engine/engine/components"

type PositionSystem interface {
	GetPosition() *components.Position
}

func ChangePosition(incX, incY int, in PositionSystem) {
	position := in.GetPosition()
	position.X += incX
	position.Y += incY
}

func CheckCollision(direction uint, inMove, inCllsn PositionSystem) int {
	incX, incY := GetIncrementMoveSystem(direction)
	pMove := inMove.GetPosition()
	pCllsn := inCllsn.GetPosition()
	if pCllsn.X == pMove.X+incX && pCllsn.Y == pMove.Y+incY {
		return Success
	}
	return Fail
}
