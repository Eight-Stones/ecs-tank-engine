package engine

import (
	"eight-stones/ecs-tank-engine/engine/common"
	"eight-stones/ecs-tank-engine/engine/pkg/utils"
	"eight-stones/ecs-tank-engine/engine/systems"
	"time"
)

func (f *Field) rotate(obj systems.CommonSystem, direction uint, now time.Time) int {
	doing := 0b0
	if !systems.CanRotate(obj, now) {
		return doing | common.FailRotate | common.Ban
	}

	rotatement := obj.(systems.RotatementSystem)
	systems.RotateMoveSystem(rotatement, direction)
	systems.SetRotateDone(rotatement, now)

	return doing | common.OkRotate
}

func (f *Field) move(obj systems.CommonSystem, now time.Time) int {
	doing := 0b0
	if !systems.CanStep(obj, now) {
		return doing | common.FailStep | common.Ban
	}

	movement := obj.(systems.MovementSystem)

	systems.SetStepDone(movement, now)

	doing = doing | f.checkBorder(movement.GetPosition().Direction, movement)
	switch {
	case utils.CheckBitMask(doing, common.Disappear):
		systems.Disappear(obj)
		return doing | common.FailStep
	case utils.CheckBitMask(doing, common.FailBorder):
		return doing | common.FailStep
	}

	for _, objPosition := range f.getAllCanPosition() {
		doing = doing | f.checkCollision(movement, objPosition)
	}

	if utils.CheckBitMask(doing, common.OkCollision) {
		return (doing ^ common.FailCollision) | common.FailStep
	}

	systems.StepMoveSystem(movement)

	doing = doing | common.OkStep

	return doing
}
