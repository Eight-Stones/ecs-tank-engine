package engine

import (
	"eight-stones/ecs-tank-engine/engine/common"
	"eight-stones/ecs-tank-engine/engine/pkg/utils"
	"time"
)

// Rotate rotates entities.Tank.
func (f *Field) Rotate(id string, direction uint) int {
	f.appInfo.mutex.Lock()
	defer f.appInfo.mutex.Unlock()
	obj, doing := f.find(id)
	if utils.CheckBitMask(doing, common.NotFound) {
		return doing
	}
	return f.rotate(obj, direction, time.Now())
}

// Move moves entities.Tank in select direction.
func (f *Field) Move(id string, direction uint) int {
	f.appInfo.mutex.Lock()
	defer f.appInfo.mutex.Unlock()
	obj, doing := f.find(id)
	if utils.CheckBitMask(doing, common.NotFound) {
		return doing
	}
	now := time.Now()

	code := f.rotate(obj, direction, now)
	if utils.CheckBitMask(code, common.FailRotate) {
		return code
	}
	return code | f.move(obj, now)
}

// Shoot shoots from entities.Tank in select direction.
func (f *Field) Shoot(id string) int {
	f.appInfo.mutex.Lock()
	defer f.appInfo.mutex.Unlock()
	obj, code := f.find(id)
	if utils.CheckBitMask(code, common.NotFound) {
		return code
	}

	return code | f.shoot(obj)
}

// Vision return small area around entities.Tank.
func (f *Field) Vision(id string) (int, [][]string) {
	f.appInfo.mutex.Lock()
	defer f.appInfo.mutex.Unlock()
	_, code := f.find(id)
	if utils.CheckBitMask(code, common.NotFound) {
		return code, nil
	}

	return code, nil
}

// Radar return big area around entities.Tank, but has recharge.
func (f *Field) Radar(id string) (int, [][]string) {
	f.appInfo.mutex.Lock()
	defer f.appInfo.mutex.Unlock()
	_, code := f.find(id)
	if utils.CheckBitMask(code, common.NotFound) {
		return code, nil
	}
	return code, nil
}
