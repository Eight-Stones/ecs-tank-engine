package engine

import (
	"context"
	"eight-stones/ecs-tank-engine/engine/config"
	"eight-stones/ecs-tank-engine/engine/entities"
	"eight-stones/ecs-tank-engine/engine/systems"
	"sync"
)

// MetaInfo inner info.
type MetaInfo struct {
	NumberGamers        int
	MaxNumberGamers     int
	SizeX               int
	SizeY               int
	PreSelectPlaces     [][]int
	PreSelectDirections []uint
}

// AppInfo inner app objects.
type AppInfo struct {
	mutex *sync.Mutex
	jobWG *sync.WaitGroup
}

// Field game field.
type Field struct {
	cfg          *config.Config
	appInfo      AppInfo
	metaInfo     MetaInfo
	NumberGamers int
	Objects      []systems.InfoSystem
	DeadObjects  []systems.InfoSystem
}

// New create new instance.
func New(cfg *config.Config) Field {
	return Field{
		cfg: cfg,
		appInfo: AppInfo{
			mutex: &sync.Mutex{},
			jobWG: &sync.WaitGroup{},
		},
		metaInfo: MetaInfo{
			MaxNumberGamers:     cfg.Game.MaxGamers,
			SizeX:               cfg.Game.SizeX,
			SizeY:               cfg.Game.SizeY,
			PreSelectPlaces:     cfg.Game.PreSelectPlaces,
			PreSelectDirections: cfg.Game.PreSelectDirection,
		},
		NumberGamers: 0,
		Objects:      nil,
		DeadObjects:  nil,
	}
}

// Start begin processes of engine.
func (f *Field) Start(ctx context.Context) {
	for idx, obj := range f.Objects {
		if tank, ok := obj.(*entities.Tank); ok {
			tank.Position.X = f.metaInfo.PreSelectPlaces[idx][0]
			tank.Position.Y = f.metaInfo.PreSelectPlaces[idx][1]
			tank.Position.Direction = f.metaInfo.PreSelectDirections[idx]
		}
	}
	f.runJobs(ctx)
}
