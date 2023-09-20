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
		}
	}
	f.runJobs(ctx)
}

// GetInfoChannel return info channel for data.
func (f *Field) GetInfoChannel() chan interface{} {
	return f.cache.getOut()
}
