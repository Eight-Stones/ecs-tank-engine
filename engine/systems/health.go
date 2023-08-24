package systems

import "eight-stones/ecs-tank-engine/engine/components"

type HealthSystem interface {
	GetHealth() *components.Health
}

func IsAliveHealthSystem(in HealthSystem) bool {
	if in.GetHealth().HitPoints <= 0 {
		return false
	}
	return true
}

func ChangeHPLevelHealthSystem(in HealthSystem, delta int) {
	health := in.GetHealth()
	switch {
	case health.HitPoints-delta <= health.MaxHitPoints:
		health.HitPoints -= delta
	case health.HitPoints-delta > health.MaxHitPoints:
		health.HitPoints = health.MaxHitPoints
	case health.HitPoints-delta <= 0:
		health.HitPoints = 0
	}
}

func Disappear(in CommonSystem) {
	obj := in.(HealthSystem)
	ChangeHPLevelHealthSystem(obj, obj.GetHealth().HitPoints)
}
