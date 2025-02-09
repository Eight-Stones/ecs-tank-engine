package engine

import (
	"time"

	"ecs-tank-engine/common"
	"ecs-tank-engine/components"
	"ecs-tank-engine/pkg/utils"
	"ecs-tank-engine/systems"
)

// Rotate rotates entities.Tank.
func (f *Field) Rotate(id string, direction components.Direction) int {
	f.sync.mutex.Lock()
	defer f.sync.mutex.Unlock()
	obj, code := f.find(id)
	if utils.CheckBitMask(code, common.NotFound) {
		return code
	}

	code = code | f.rotate(obj, direction, time.Now())
	systems.AddAction(code, obj)

	return code
}

// Move moves entities.Tank in select direction.
func (f *Field) Move(id string, direction components.Direction) int {
	f.sync.mutex.Lock()
	defer f.sync.mutex.Unlock()
	obj, doing := f.find(id)
	if utils.CheckBitMask(doing, common.NotFound) {
		return doing
	}
	now := time.Now()

	code := f.rotate(obj, direction, now)
	if utils.CheckBitMask(code, common.FailRotate) {
		return code
	}

	code = code | f.move(obj, now)

	systems.AddAction(code, obj)

	return code
}

// Shoot shoots from entities.Tank in select direction.
func (f *Field) Shoot(id string) int {
	f.sync.mutex.Lock()
	defer f.sync.mutex.Unlock()
	obj, code := f.find(id)
	if utils.CheckBitMask(code, common.NotFound) {
		return code
	}

	code = code | f.shoot(obj)
	systems.AddAction(code, obj)

	return code
}

// Vision return small area around entities.Tank.
func (f *Field) Vision(id string) (int, View) {
	f.sync.mutex.Lock()
	defer f.sync.mutex.Unlock()
	obj, code := f.find(id)
	if utils.CheckBitMask(code, common.NotFound) {
		return code, nil
	}

	codeView, view := f.vision(obj)
	code = code | codeView

	systems.AddAction(code, obj)

	return code, view
}

// Radar return big area around entities.Tank, but has recharge.
func (f *Field) Radar(id string) (int, View) {
	f.sync.mutex.Lock()
	defer f.sync.mutex.Unlock()
	obj, code := f.find(id)
	if utils.CheckBitMask(code, common.NotFound) {
		return code, nil
	}

	codeView, view := f.radar(obj)
	code = code | codeView

	systems.AddAction(code, obj)

	return code, view
}
