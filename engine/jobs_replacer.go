package engine

import (
	"context"
	"eight-stones/ecs-tank-engine/engine/systems"
	"time"
)

// autoReplaceDeadJob start process of replacing dead object to other list.
func (f *Field) autoReplaceDeadJob(ctx context.Context) {
	defer f.sync.jobWG.Done()
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
	f.sync.mutex.Lock()
	defer f.sync.mutex.Unlock()
	f.autoReplace()
}

// autoReplaceDead replace dead object to other list.
func (f *Field) autoReplace() {
	alive := make([]systems.InfoSystem, 0, len(f.Objects))
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
		f.cache.saveRemove(f.Objects[idx].GetInfo().Id, f.Objects[idx].GetInfo().Type)
	}
	f.Objects = alive
}

// replaceDeadById replace entity from alive to dead by id.
func (f *Field) replaceDeadById(id string) {
	var selectIDx int
	for idx := range f.Objects {
		if f.Objects[idx].GetInfo().Id == id {
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
	f.cache.saveRemove(f.Objects[selectIDx].GetInfo().Id, f.Objects[selectIDx].GetInfo().Type)
	f.Objects = append(f.Objects[:selectIDx], f.Objects[selectIDx+1:]...)
}
