package engine

import (
	"context"
	"eight-stones/ecs-tank-engine/engine/entities"
)

// Start begin processes of engine.
func (f *Field) Start(ctx context.Context) {
	for idx, obj := range f.Objects {
		if tank, ok := obj.(*entities.Tank); ok {
			tank.Position.X = f.gameInfo.PreSelectPlaces[idx][0]
			tank.Position.Y = f.gameInfo.PreSelectPlaces[idx][1]
			tank.Position.Direction = f.gameInfo.PreSelectDirections[idx]
			f.cache.saveCreate(
				tank.Id,
				tank.Info.Type,
				tank.Position.Direction,
				[]int{tank.Position.X, tank.Position.Y},
				tank.Health.HitPoints,
			)
		}
	}
	ctx, cancel := context.WithCancel(ctx)
	f.cancelCtx = cancel
	f.runJobs(ctx)
}

// Stop ends processes of engine.
func (f *Field) Stop() {
	f.cancelCtx()
	f.sync.jobWG.Wait()
}

// GetInfoChannel return info channel for data.
func (f *Field) GetInfoChannel() chan Info {
	return f.cache.getOut()
}
