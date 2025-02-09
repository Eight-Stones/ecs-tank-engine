package engine

import (
	"github.com/Eight-Stones/ecs-tank-engine/common"
	"github.com/Eight-Stones/ecs-tank-engine/systems"
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

// makeDamage makes damage to first from second.
func (f *Field) makeDamage(first, second systems.InfoSystem) int {
	fhs := first.(systems.HealthSystem)
	sds := second.(systems.DamageSystem)
	oldHitPoints := fhs.GetHealth().HitPoints
	systems.CauseHitDamageSystem(fhs, sds)
	newHitPoints := fhs.GetHealth().HitPoints
	f.cache.saveCollision(
		first.GetInfo().Id,
		first.GetInfo().Type,
		oldHitPoints,
		newHitPoints,
	)
	return common.Damaged
}
