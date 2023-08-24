package engine

import (
	"context"
	"eight-stones/ecs-tank-engine/engine/systems"
	"time"
)

// autoReplaceDeadJob автоматически переносит мертвые игровые объекты в лог объектов.
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

// autoReplaceDead переносит мертвые игровые объекты в лог объектов.
func (f *Field) autoReplaceDead(_ context.Context) {
	f.appInfo.mutex.Lock()
	defer f.appInfo.mutex.Unlock()
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
