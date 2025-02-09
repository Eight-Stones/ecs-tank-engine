package systems

import (
	"ecs-tank-engine/components"
	"time"
)

// MovementSystem describes movement system.
type MovementSystem interface {
	GetMovement() *components.Movement
	PositionSystem
}

// CanStep checks that system can do step.
func CanStep(in InfoSystem, now time.Time) bool {
	movement, ok := in.(MovementSystem)
	if !ok {
		return false
	}
	return movement.GetMovement().Recharge.IsRechargeDone(now)
}

// SetStepDone set actions 'DoStep' as success done and sets recharge since time.
func SetStepDone(in MovementSystem, now time.Time) {
	recharge := in.GetMovement().Recharge
	recharge.SetUntil(now)
}

// DoStep make step in current direction.
func DoStep(in MovementSystem) {
	incX, incY := GetIncrementPosition(in.GetPosition().Direction)
	ChangePosition(incX, incY, in)
}
