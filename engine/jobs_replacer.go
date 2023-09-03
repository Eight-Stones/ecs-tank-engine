package engine

import (
	"context"
	"eight-stones/ecs-tank-engine/engine/systems"
	"time"
)

// autoReplaceDeadJob start process of replacing dead object to other list.
func (f *Field) autoReplaceDeadJob(ctx context.Context) {
	defer f.appInfo.jobWG.Done()
	ticker := time.NewTicker(f.cfg.Game.Jobs.Replacer)
	for {
		select {
		case <-ticker.C:
			f.autoReplaceDead(ctx)
		case <-ctx.Done():
			return
		}
	}
}

// autoReplaceDead replace dead object to other list.
func (f *Field) autoReplaceDead(_ context.Context) {
	f.appInfo.mutex.Lock()
	defer f.appInfo.mutex.Unlock()
	f.autoReplace()
}

// autoReplaceDead replace dead object to other list.
func (f *Field) autoReplace() {
	alive := make([]systems.CommonSystem, 0, len(f.Objects))
	for idx := range f.Objects {
		healthObject, ok := f.Objects[idx].(systems.HealthSystem)
		if !ok {
			alive = append(alive, f.Objects[idx])
			continue
		}

		if systems.IsAliveHealthSystem(healthObject) {
			alive = append(alive, f.Objects[idx])
			continue
		}

		f.DeadObjects = append(f.DeadObjects, f.Objects[idx])
	}
	f.Objects = alive
}

func (f *Field) replaceDeadById(id string) {
	var selectIDx int
	for idx := range f.Objects {
		if f.Objects[idx].GetCommon().Id == id {
			selectIDx = idx
			break
		}
	}

	if selectIDx == 0 {
		return
	}

	healthObject, ok := f.Objects[selectIDx].(systems.HealthSystem)
	if !ok {
		return
	}

	if systems.IsAliveHealthSystem(healthObject) {
		return
	}

	f.DeadObjects = append(f.DeadObjects, f.Objects[selectIDx])
	f.Objects = append(f.Objects[:selectIDx], f.Objects[selectIDx+1:]...)
}
