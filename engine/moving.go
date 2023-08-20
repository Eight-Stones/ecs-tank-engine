package engine

import (
	"eight-stones/ecs-tank-engine/engine/common"
	"eight-stones/ecs-tank-engine/engine/pkg"
	"eight-stones/ecs-tank-engine/engine/systems"
	"time"
)

func (f *Field) rotate(id string, direction uint) int {
	tank, code := f.find(id)
	doing := 0b0 | code
	if pkg.CheckBitMask(code, common.NotFound) {
		return doing | common.FailRotate
	}

	now := time.Now()
	if !systems.CanRotate(tank, now) {
		return doing | common.FailRotate
	}

	systems.RotateMoveSystem(direction, tank)
	systems.SetRotateDone(tank, now)

	return doing | common.OkRotate
}

func (f *Field) move(id string) int {
	tank, code := f.find(id)
	doing := 0b0 | code

	if pkg.CheckBitMask(doing, common.NotFound) {
		return doing | common.FailStep
	}

	now := time.Now()
	if !systems.CanStep(tank, now) {
		return doing | common.FailStep
	}

	systems.SetStepDone(tank, now)

	doing = doing | f.checkBorder(tank.Direction, tank)
	if pkg.CheckBitMask(doing, common.FailBorder) {
		return doing | common.FailStep
	}

	canPositionObjects := f.getAllCanPosition()
	for _, obj := range canPositionObjects {
		doing = doing | f.checkCollision(tank, obj)
	}

	if pkg.CheckBitMask(doing, common.OkCollision) {
		return doing | common.FailStep
	}

	systems.StepMoveSystem(tank)

	doing = doing | common.OkStep

	return doing
}
