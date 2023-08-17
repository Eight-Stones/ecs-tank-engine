package systems

import (
	"eight-stones/ecs-tank-engine/engine/common"
	"eight-stones/ecs-tank-engine/engine/components"
)

type MovementSystem interface {
	GetMovement() *components.Movement
	PositionSystem
}

type RotatementSystem interface {
	GetMovement() *components.Movement
	CanRotate()
	PositionSystem
}

type NotInterruptMovementSystem interface {
	DoesNotInterruptMovement()
}

type AutoMovementSystem interface {
	CanAutoMove()
	MovementSystem
}

func GetIncrementMoveSystem(direction uint) (x, y int) {
	values := common.MovementValue[direction]
	return values[0], values[1]
}

func CheckCollision(first MovementSystem, second PositionSystem) int {
	if first == second {
		return Fail
	}

	incX, incY := GetIncrementMoveSystem(first.GetMovement().Direction)
	pMove := first.GetPosition()
	pCllsn := second.GetPosition()

	if pCllsn.X == pMove.X+incX && pCllsn.Y == pMove.Y+incY {
		return Success
	}
	return Fail
}

func RotateMoveSystem(direction uint, in RotatementSystem) {
	in.GetMovement().Direction = direction
}

func StepMoveSystem(in MovementSystem) {
	incX, incY := GetIncrementMoveSystem(in.GetMovement().Direction)
	ChangePosition(incX, incY, in)
}
