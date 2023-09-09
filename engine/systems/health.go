package systems

import "eight-stones/ecs-tank-engine/engine/components"

// HealthSystem describes Health system.
type HealthSystem interface {
	GetHealth() *components.Health
}

// IsAliveHealthSystem checks whether health is sufficient to continue living.
func IsAliveHealthSystem(in HealthSystem) bool {
	if in.GetHealth().HitPoints <= 0 {
		return false
	}
	return true
}

// ChangeHPLevelHealthSystem changes the health level to the delta indicator.
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

// Disappear resets the number of lives to zero.
func Disappear(in InfoSystem) {
	obj := in.(HealthSystem)
	ChangeHPLevelHealthSystem(obj, obj.GetHealth().HitPoints)
}
