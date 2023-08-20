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
	recharge := in.GetRotatement().Recharge
	if recharge.FreeAction != 0 && recharge.Until.After(now) {
		return true
	}
	return false
}

func SetRotateDone(in RotatementSystem, now time.Time) {
	recharge := in.GetRotatement().Recharge
	recharge.DecFreeAction()
	recharge.SetUntil(now)

}

func RotateMoveSystem(direction uint, in RotatementSystem) int {
	if !CanRotate(in, time.Now()) {
		return Fail
	}
	in.GetMovement().Direction = direction

	return Success
}

func StepMoveSystem(in MovementSystem) {
	incX, incY := GetIncrementMoveSystem(in.GetMovement().Direction)
	ChangePosition(incX, incY, in)
}
