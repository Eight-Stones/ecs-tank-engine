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
	return f.rotate(id, direction, time.Now())
}

// Move moves entities.Tank in select direction.
func (f *Field) Move(id string, direction uint) int {
	f.appInfo.mutex.Lock()
	defer f.appInfo.mutex.Unlock()
	now := time.Now()
	code := f.rotate(id, direction, now)

	if utils.CheckBitMask(code, common.FailRotate) {
		return code
	}
	return code | f.move(id, now)
}

// Shoot shoots from entities.Tank in select direction.
func (f *Field) Shoot(id string) int {
	f.appInfo.mutex.Lock()
	defer f.appInfo.mutex.Unlock()
	return f.shoot(id)
}

// Vision return small area around entities.Tank.
func (f *Field) Vision(id string) (int, [][]string) {
	f.appInfo.mutex.Lock()
	defer f.appInfo.mutex.Unlock()
	return -1, nil
}

// Radar return big area around entities.Tank, but has recharge.
func (f *Field) Radar(id string) (int, [][]string) {
	f.appInfo.mutex.Lock()
	defer f.appInfo.mutex.Unlock()
	return -1, nil
}
