package engine

import (
	"eight-stones/ecs-tank-engine/engine/common"
	"eight-stones/ecs-tank-engine/engine/components"
	"eight-stones/ecs-tank-engine/engine/pkg/utils"
	"eight-stones/ecs-tank-engine/engine/systems"
	"time"
)

// rotate rotates select object.
func (f *Field) rotate(obj systems.InfoSystem, direction components.Direction, now time.Time) int {
	doing := 0b0
	if !systems.CanRotate(obj, now) {
		return doing | common.FailRotate | common.Ban
	}

	rotatement := obj.(systems.RotatementSystem)
	systems.DoRotate(rotatement, direction)
	systems.SetRotateDone(rotatement, now)

	f.cache.saveRotatement(
		obj.GetInfo().Id,
		obj.GetInfo().Type,
		rotatement.GetPosition().Direction,
		direction,
	)

	return doing | common.OkRotate
}

// move moves select object by his direction.
func (f *Field) move(obj systems.InfoSystem, now time.Time) int {
	doing := 0b0
	if !systems.CanStep(obj, now) {
		return doing | common.FailStep | common.Ban
	}

	movement := obj.(systems.MovementSystem)

	systems.SetStepDone(movement, now)

	doing = doing | f.checkBorder(movement.GetPosition().Direction, movement)
	switch {
	case utils.CheckBitMask(doing, common.Disappear):
		systems.Disappear(obj.(systems.HealthSystem))
		return doing | common.FailStep
	case utils.CheckBitMask(doing, common.FailBorder):
		return doing | common.FailStep
	}

	for _, objPosition := range f.getAllCanPosition() {
		doing = doing | f.checkCollision(movement, objPosition)
	}

	if utils.CheckBitMask(doing, common.OkCollision) {
		return (doing ^ common.NoCollision) | common.FailStep
	}

	oldP := []int{movement.GetPosition().X, movement.GetPosition().Y}

	systems.DoStep(movement)

	newP := []int{movement.GetPosition().X, movement.GetPosition().Y}

	f.cache.saveStep(
		obj.GetInfo().Id,
		obj.GetInfo().Type,
		oldP,
		newP,
	)

	return doing | common.OkStep
}
