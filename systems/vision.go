package systems

import (
	"time"

	"github.com/Eight-Stones/ecs-tank-engine/components"
)

// VisionSystem describes vision system.
type VisionSystem interface {
	GetVision() *components.Vision
	PositionSystem
}

// CanVision checks that system can do vision.
func CanVision(in InfoSystem, now time.Time) bool {
	vision, ok := in.(VisionSystem)
	if !ok {
		return false
	}
	return vision.GetVision().Recharge.IsRechargeDone(now)
}

// SetVisionDone set actions 'Vision' as success done and sets recharge since time.
func SetVisionDone(in VisionSystem, now time.Time) {
	recharge := in.GetVision().Recharge
	recharge.SetUntil(now)
}
