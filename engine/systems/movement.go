package systems

import (
	"eight-stones/ecs-tank-engine/engine/common"
	"eight-stones/ecs-tank-engine/engine/components"
	"time"
)

type MovementSystem interface {
	GetMovement() *components.Movement
	PositionSystem
}

type RotatementSystem interface {
	GetRotatement() *components.Rotatement
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

func CheckCollision(first PositionSystem, second PositionSystem) int {
	if first == second {
		return Fail
	}

	incX, incY := GetIncrementMoveSystem(first.GetPosition().Direction)
	pMove := first.GetPosition()
	pCllsn := second.GetPosition()

	if pCllsn.X == pMove.X+incX && pCllsn.Y == pMove.Y+incY {
		return Success
	}
	return Fail
}

func CanRotate(in CommonSystem, now time.Time) bool {
	rotatement, ok := in.(RotatementSystem)
	if !ok {
		return false
	}
	return rotatement.GetRotatement().Recharge.IsRechargeDone(now)
}

func SetRotateDone(in RotatementSystem, now time.Time) {
	recharge := in.GetRotatement().Recharge
	recharge.SetUntil(now)
}

func RotateMoveSystem(in RotatementSystem, direction uint) {
	in.GetPosition().Direction = direction
}

func CanStep(in CommonSystem, now time.Time) bool {
	movement, ok := in.(MovementSystem)
	if !ok {
		return false
	}
	return movement.GetMovement().Recharge.IsRechargeDone(now)
}

func SetStepDone(in MovementSystem, now time.Time) {
	recharge := in.GetMovement().Recharge
	recharge.SetUntil(now)
}

func StepMoveSystem(in MovementSystem) {
	incX, incY := GetIncrementMoveSystem(in.GetPosition().Direction)
	ChangePosition(incX, incY, in)
}
