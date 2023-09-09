package systems

import "eight-stones/ecs-tank-engine/engine/components"

// DamageSystem describes interface of Damage system.
type DamageSystem interface {
	GetDamage() *components.Damage
	HealthSystem
}

// CauseHitDamageSystem deals health damage based on components.
func CauseHitDamageSystem(damageTaker HealthSystem, damageDealer DamageSystem) {
	ChangeHPLevelHealthSystem(damageTaker, damageDealer.GetDamage().DamagePoints)
}
