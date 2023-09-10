package systems

import (
	"eight-stones/ecs-tank-engine/engine/components"
	"time"
)

// RadarSystem describes radar system.
type RadarSystem interface {
	GetRadar() *components.Radar
	PositionSystem
}

// CanRadar checks that system can do radar.
func CanRadar(in InfoSystem, now time.Time) bool {
	radar, ok := in.(RadarSystem)
	if !ok {
		return false
	}
	return radar.GetRadar().Recharge.IsRechargeDone(now)
}

// SetRadarDone set actions 'Radar' as success done and sets recharge since time.
func SetRadarDone(in RadarSystem, now time.Time) {
	recharge := in.GetRadar().Recharge
	recharge.SetUntil(now)
}
