package engine

import (
	"eight-stones/ecs-tank-engine/engine/common"
	"eight-stones/ecs-tank-engine/engine/systems"
)

// defineDamageType defines type of damage on object.
func (f *Field) defineDamageType(in systems.InfoSystem) int {
	result := common.DoNothing
	if _, ok := in.(systems.HealthSystem); ok {
		result = result | common.CanOnlyDamaged
	}
	if _, ok := in.(systems.DamageSystem); ok {
		result = result | common.CanOnlyDamage
	}
	return result
}

// makeDamage make damage to first from second.
func (f *Field) makeDamage(first, second systems.InfoSystem) int {
	systems.CauseHitDamageSystem(first.(systems.HealthSystem), second.(systems.DamageSystem))
	return common.Damaged
}
