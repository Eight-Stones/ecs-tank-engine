package engine

import (
	"eight-stones/ecs-tank-engine/engine/common"
	"eight-stones/ecs-tank-engine/engine/systems"
)

// checkBorder check collision with border of map.
func (f *Field) checkBorder(direction uint, in systems.PositionSystem) int {
	incX, incY := systems.GetIncrementMoveSystem(direction)
	object := in.GetPosition()

	breakLeft := object.X+incX < 0
	breakRight := object.X+incX > f.metaInfo.SizeX-1
	breakBottom := object.Y+incY < 0
	breakTop := object.Y+incY > f.metaInfo.SizeY-1

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
	if systems.CheckCollision(first, second) == systems.Fail {
		return common.NoCollision
	}

	if f.makeCollision(first.(systems.CommonSystem), second.(systems.CommonSystem)) == common.DoNothing {
		return common.OkCollision
	}

	if _, ok := second.(systems.NotInterruptMovementSystem); ok {
		return common.NotInterruptOkCollision
	}

	return common.OkCollision

}

// makeCollision process the collision between objects.
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
