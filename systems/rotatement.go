package systems

import (
	"time"

	"github.com/Eight-Stones/ecs-tank-engine/v2/components"
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
func DoRotate(in RotatementSystem, direction components.Direction) {
	// todo need to check direction on existing
	in.GetPosition().Direction = direction
}
