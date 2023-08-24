package systems

import "eight-stones/ecs-tank-engine/engine/components"

type DamageSystem interface {
	GetDamage() *components.Damage
	HealthSystem
}

func CauseHitDamageSystem(damageTaker HealthSystem, damageDealer DamageSystem) {
	ChangeHPLevelHealthSystem(damageTaker, damageDealer.GetDamage().DamagePoints)
}
