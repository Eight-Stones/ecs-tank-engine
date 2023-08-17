package engine

import (
	"eight-stones/ecs-tank-engine/engine/common"
	"eight-stones/ecs-tank-engine/engine/systems"
)

func (f *Field) checkBorder(direction uint, in systems.PositionSystem) int {
	incX, incY := systems.GetIncrementMoveSystem(direction)
	object := in.GetPosition()
	if object.X+incX < 0 || object.X+incX > f.Border.X || object.Y+incY < 0 || object.Y+incY > f.Border.Y {
		return common.BreakBorder
	}
	return common.NoneBreakBorder
}

func (f *Field) checkCollision(first systems.MovementSystem, second systems.PositionSystem) int {
	if systems.CheckCollision(first, second) == systems.Success {
		code := f.makeCollision(first.(systems.CommonSystem), second.(systems.CommonSystem))
		if code == common.DoNothing {
			return common.CollisionSuccess
		}

		if _, ok := second.(systems.NotInterruptMovementSystem); ok {
			return common.CollisionSuccessNotInterruptMove
		}

		return common.CollisionSuccess
	}

	return common.NoneCollision
}

func (f *Field) makeCollision(first, second systems.CommonSystem) int {
	fInfo, sInfo := f.defineDamageType(first), f.defineDamageType(second)
	code := common.DoNothing
	if fInfo == common.DoNothing || sInfo == common.DoNothing {
		return code
	}

	bothCanDamageAndDamaged := fInfo&common.CanDamagedAndDamage == common.CanDamagedAndDamage &&
		sInfo&common.CanDamagedAndDamage == common.CanDamagedAndDamage

	firstAllSecondOnlyDamaged := fInfo&common.CanDamagedAndDamage == common.CanDamagedAndDamage &&
		sInfo&common.CanOnlyDamaged == common.CanOnlyDamaged

	firstOnlyDamagedSecondAll := fInfo&common.CanOnlyDamaged == common.CanOnlyDamaged &&
		sInfo&common.CanDamagedAndDamage == common.CanDamagedAndDamage

	switch {
	case bothCanDamageAndDamaged:
		f.makeDamage(second, first)
		f.makeDamage(first, second)
		code = code | common.BothDamaged
	case firstAllSecondOnlyDamaged:
		f.makeDamage(second, first)
		code = code | common.Damaged
	case firstOnlyDamagedSecondAll:
		f.makeDamage(first, second)
		code = code | common.Damaged
	}

	return code
}
