package engine

import (
	"ecs-tank-engine/common"
	"ecs-tank-engine/components"
	"ecs-tank-engine/systems"
)

// checkBorder check collision with border of map.
func (f *Field) checkBorder(direction components.Direction, in systems.PositionSystem) int {
	incX, incY := systems.GetIncrementPosition(direction)
	object := in.GetPosition()

	breakLeft := object.X+incX < 0
	breakRight := object.X+incX > f.gameInfo.SizeX-1
	breakBottom := object.Y+incY < 0
	breakTop := object.Y+incY > f.gameInfo.SizeY-1

	isBreakBorder := breakLeft || breakRight || breakBottom || breakTop
	_, isCanDisappear := in.(systems.AutoMovementSystem)

	switch {
	case isBreakBorder && isCanDisappear:
		return common.FailBorder | common.Disappear
	case isBreakBorder:
		return common.FailBorder

	}

	return common.OkBorder
}

// checkCollision check collision with other object.
func (f *Field) checkCollision(first systems.MovementSystem, second systems.PositionSystem) int {
	if systems.IsCollision(first, second) == systems.Fail {
		return common.NoCollision
	}

	if f.makeCollision(first.(systems.InfoSystem), second.(systems.InfoSystem)) == common.DoNothing {
		return common.OkCollision
	}

	if _, ok := second.(systems.NotInterruptMovementSystem); ok {
		return common.NotInterruptOkCollision
	}

	return common.OkCollision

}

// makeCollision process the collision between objects.
func (f *Field) makeCollision(first, second systems.InfoSystem) int {
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
