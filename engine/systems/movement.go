package systems

import (
	"eight-stones/ecs-tank-engine/engine/common"
	"eight-stones/ecs-tank-engine/engine/components"
)

type MovementSystem interface {
	GetMovement() *components.Movement
	PositionSystem
}

type AutoMovementSystem interface {
	CanAutoMove()
	MovementSystem
}

func GetIncrementMoveSystem(direction uint) (x, y int) {
	values := common.MovementValue[direction]
	return values[0], values[1]
}

func StepMoveSystem(direction uint, in MovementSystem) {
	incX, incY := GetIncrementMoveSystem(direction)
	in.GetMovement().Direction = direction
	ChangePosition(incX, incY, in)
}

func AutoStepMoveSystem(in AutoMovementSystem) {
	incX, incY := GetIncrementMoveSystem(in.GetMovement().Direction)
	ChangePosition(incX, incY, in.GetPosition())
}
