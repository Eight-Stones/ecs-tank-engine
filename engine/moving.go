package engine

import (
	"eight-stones/ecs-tank-engine/engine/common"
	"eight-stones/ecs-tank-engine/engine/pkg/utils"
	"eight-stones/ecs-tank-engine/engine/systems"
	"time"
)

func (f *Field) rotate(id string, direction uint) int {
	tank, code := f.find(id)
	doing := 0b0 | code
	if utils.CheckBitMask(code, common.NotFound) {
		return doing | common.FailRotate
	}

	now := time.Now()
	if !systems.CanRotate(tank, now) {
		return doing | common.FailRotate | common.Ban
	}

	systems.RotateMoveSystem(direction, tank)
	systems.SetRotateDone(tank, now)

	return doing | common.OkRotate
}

func (f *Field) move(id string) int {
	tank, code := f.find(id)
	doing := 0b0 | code

	if utils.CheckBitMask(doing, common.NotFound) {
		return doing | common.FailStep
	}

	now := time.Now()
	if !systems.CanStep(tank, now) {
		return doing | common.FailStep | common.Ban
	}

	systems.SetStepDone(tank, now)

	doing = doing | f.checkBorder(tank.Direction, tank)
	if utils.CheckBitMask(doing, common.FailBorder) {
		return doing | common.FailStep
	}

	for _, obj := range f.getAllCanPosition() {
		doing = doing | f.checkCollision(tank, obj)
	}

	if utils.CheckBitMask(doing, common.OkCollision) {
		return (doing ^ common.FailCollision) | common.FailStep
	}

	systems.StepMoveSystem(tank)

	doing = doing | common.OkStep

	return doing
}
