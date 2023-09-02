package engine

import (
	"eight-stones/ecs-tank-engine/engine/common"
	"eight-stones/ecs-tank-engine/engine/pkg/utils"
	"eight-stones/ecs-tank-engine/engine/systems"
	"time"
)

// Rotate rotates entities.Tank.
func (f *Field) Rotate(id string, direction uint) int {
	f.appInfo.mutex.Lock()
	defer f.appInfo.mutex.Unlock()
	obj, code := f.find(id)
	if utils.CheckBitMask(code, common.NotFound) {
		return code
	}

	code = code | f.rotate(obj, direction, time.Now())
	systems.AddActionStatisticSystem(code, obj)

	return code
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

	code = code | f.move(obj, now)
	systems.AddActionStatisticSystem(code, obj)

	return code
}

// Shoot shoots from entities.Tank in select direction.
func (f *Field) Shoot(id string) int {
	f.appInfo.mutex.Lock()
	defer f.appInfo.mutex.Unlock()
	obj, code := f.find(id)
	if utils.CheckBitMask(code, common.NotFound) {
		return code
	}

	code = code | f.shoot(obj)
	systems.AddActionStatisticSystem(code, obj)

	return code
}

// Vision return small area around entities.Tank.
func (f *Field) Vision(id string) (int, [][]string) {
	f.appInfo.mutex.Lock()
	defer f.appInfo.mutex.Unlock()
	_, code := f.find(id)
	if utils.CheckBitMask(code, common.NotFound) {
		return code, nil
	}

	systems.AddActionStatisticSystem(code, nil)

	return code, nil
}

func (f *Field) vision() {

}

// Radar return big area around entities.Tank, but has recharge.
func (f *Field) Radar(id string) (int, [][]string) {
	f.appInfo.mutex.Lock()
	defer f.appInfo.mutex.Unlock()
	_, code := f.find(id)
	if utils.CheckBitMask(code, common.NotFound) {
		return code, nil
	}

	systems.AddActionStatisticSystem(code, nil)

	return code, nil
}
