package systems

import (
	"eight-stones/ecs-tank-engine/engine/components"
	"time"
)

type ShootingSystem interface {
	GetShooting() *components.Shooting
}

func CanShoot(in CommonSystem, now time.Time) bool {
	shooting, ok := in.(ShootingSystem)
	if !ok {
		return false
	}
	return shooting.GetShooting().Recharge.IsRechargeDone(now) && shooting.GetShooting().Ammo > 0
}

func SetShotDone(in ShootingSystem, now time.Time) {
	in.GetShooting().Ammo--
	in.GetShooting().Recharge.SetUntil(now)
}
