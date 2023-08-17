package engine

import (
	"eight-stones/ecs-tank-engine/engine/common"
	"eight-stones/ecs-tank-engine/engine/systems"
)

func (f *Field) rotate(id string, direction uint) int {
	tank, code := f.find(id)
	doing := 0b0 | code
	if code&common.TankNotFound == common.TankNotFound {
		return doing
	}

	systems.RotateMoveSystem(direction, tank)

	return doing
}

func (f *Field) move(id string) int {
	tank, code := f.find(id)
	doing := 0b0 | code
	if code&common.TankNotFound == common.TankNotFound {
		return doing
	}

	doing = doing | f.checkBorder(tank.Direction, tank)
	if doing&common.BreakBorder == common.BreakBorder {
		return doing
	}

	canPositionObjects := f.getAllCanPosition()
	for _, obj := range canPositionObjects {
		doing = doing | f.checkCollision(tank, obj)
	}

	if doing&common.CollisionSuccess == common.CollisionSuccess {
		return doing
	}

	systems.StepMoveSystem(tank)

	doing = doing | common.MoveSuccess

	return doing
}
