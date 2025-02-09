package systems

import (
	"time"

	"github.com/Eight-Stones/ecs-tank-engine/components"
)

// ShootingSystem describe shooting system.
type ShootingSystem interface {
	GetShooting() *components.Shooting
}

// CanShoot checks that system can do shoot.
func CanShoot(in InfoSystem, now time.Time) bool {
	shooting, ok := in.(ShootingSystem)
	if !ok {
		return false
	}
	return shooting.GetShooting().Recharge.IsRechargeDone(now) && shooting.GetShooting().Ammo > 0
}

// SetShotDone set actions 'DoShoot' as success done, dec ammo and sets recharge since time.
func SetShotDone(in ShootingSystem, now time.Time) {
	in.GetShooting().Ammo--
	in.GetShooting().Recharge.SetUntil(now)
}
