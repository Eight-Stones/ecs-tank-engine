package systems

import (
	"eight-stones/ecs-tank-engine/engine/common"
	"eight-stones/ecs-tank-engine/engine/components"
	"time"
)

// RotatementSystem describes rotatement system.
type RotatementSystem interface {
	GetRotatement() *components.Rotatement
	PositionSystem
}

// CanRotate checks that system can do rotate.
func CanRotate(in InfoSystem, now time.Time) bool {
	rotatement, ok := in.(RotatementSystem)
	if !ok {
		return false
	}
	return rotatement.GetRotatement().Recharge.IsRechargeDone(now)
}

// SetRotateDone set actions 'DoRotate' as success done and sets recharge since time.
func SetRotateDone(in RotatementSystem, now time.Time) {
	recharge := in.GetRotatement().Recharge
	recharge.SetUntil(now)
}

// DoRotate change direction.
func DoRotate(in RotatementSystem, direction common.Direction) {
	in.GetPosition().Direction = direction
}
