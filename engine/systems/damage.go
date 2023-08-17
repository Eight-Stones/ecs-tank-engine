package systems

import "eight-stones/ecs-tank-engine/engine/components"

type DamageSystem interface {
	GetDamage() *components.Damage
	HealthSystem
}

type ShootingSystem interface{}

func CauseHitDamageSystem(damageTaker HealthSystem, damageDealer DamageSystem) {
	ChangeHPLevelHealthSystem(damageTaker, damageDealer.GetDamage().DamagePoints)
}
