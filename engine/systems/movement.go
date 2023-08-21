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
	GetMovement() *components.Movement
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

func CanRotate(in RotatementSystem, now time.Time) bool {
	return in.GetRotatement().Recharge.IsRechargeDone(now)
}

func SetRotateDone(in RotatementSystem, now time.Time) {
	recharge := in.GetRotatement().Recharge
	recharge.SetUntil(now)
}

func RotateMoveSystem(direction uint, in RotatementSystem) {
	in.GetMovement().Direction = direction
}

func CanStep(in MovementSystem, now time.Time) bool {
	return in.GetMovement().Recharge.IsRechargeDone(now)
}

func SetStepDone(in MovementSystem, now time.Time) {
	recharge := in.GetMovement().Recharge
	recharge.SetUntil(now)
}

func StepMoveSystem(in MovementSystem) {
	incX, incY := GetIncrementMoveSystem(in.GetMovement().Direction)
	ChangePosition(incX, incY, in)
}
