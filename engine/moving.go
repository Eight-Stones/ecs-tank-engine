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
		return doing | common.Rotate | common.Fail
	}

	now := time.Now()
	if !systems.CanRotate(tank, now) {
		return doing | common.Rotate | common.Fail
	}

	systems.RotateMoveSystem(direction, tank)
	systems.SetRotateDone(tank, now)

	return doing | common.Rotate | common.Success
}

func (f *Field) move(id string) int {
	tank, code := f.find(id)
	doing := 0b0 | code

	if pkg.CheckBitMask(code, common.NotFound) {
		return doing | common.Step | common.Fail
	}

	now := time.Now()
	if !systems.CanStep(tank, now) {
		return doing | common.Step | common.Fail
	}

	systems.SetStepDone(tank, now)

	doing = doing | f.checkBorder(tank.Direction, tank)
	if pkg.CheckBitMask(code, common.Fail, common.Border) {
		return doing | common.Step | common.Fail
	}

	canPositionObjects := f.getAllCanPosition()
	for _, obj := range canPositionObjects {
		doing = doing | f.checkCollision(tank, obj)
	}

	if pkg.CheckBitMask(doing, common.Collision, common.Success) {
		return doing | common.Step | common.Fail
	}

	systems.StepMoveSystem(tank)

	doing = doing | common.Step | common.Success

	return doing
}
