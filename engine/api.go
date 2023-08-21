package engine

import (
	"eight-stones/ecs-tank-engine/engine/common"
	"eight-stones/ecs-tank-engine/engine/pkg/utils"
)

// Rotate rotates entities.Tank.
func (f *Field) Rotate(id string, direction uint) int {
	return f.rotate(id, direction)
}

// Move moves entities.Tank in select direction.
func (f *Field) Move(id string, direction uint) int {
	code := f.rotate(id, direction)

	if utils.CheckBitMask(code, common.FailRotate) {
		return code
	}
	return code | f.move(id)
}

// Shoot shoots from entities.Tank in select direction.
func (f *Field) Shoot(id string) int {
	return -1
}

// Vision return small area around entities.Tank.
func (f *Field) Vision(id string) int {
	return -1
}

// Radar return big area around entities.Tank, but has recharge.
func (f *Field) Radar(id string) int {
	return -1
}
