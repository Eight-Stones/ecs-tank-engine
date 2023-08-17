package engine

import (
	"eight-stones/ecs-tank-engine/engine/common"
	"eight-stones/ecs-tank-engine/engine/systems"
)

func (f *Field) defineDamageType(in systems.CommonSystem) int {
	result := common.DoNothing
	if _, ok := in.(systems.HealthSystem); ok {
		result = result | common.CanOnlyDamaged
	}
	if _, ok := in.(systems.DamageSystem); ok {
		result = result | common.CanOnlyDamage
	}
	return result
}

func (f *Field) makeDamage(first, second systems.CommonSystem) int {
	systems.CauseHitDamageSystem(first.(systems.HealthSystem), second.(systems.DamageSystem))
	return common.Damaged
}
