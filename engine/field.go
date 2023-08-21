package engine

import (
	"context"
	"eight-stones/ecs-tank-engine/engine/config"
	"eight-stones/ecs-tank-engine/engine/entities"
	"eight-stones/ecs-tank-engine/engine/systems"
	"sync"
)

// MetaInfo метаинформация.
type MetaInfo struct {
	NumberGamers    int
	MaxNumberGamers int
	SizeX           int
	SizeY           int
	PreSelectPlaces [][]int
}

type AppInfo struct {
	mutex *sync.Mutex
	jobWG *sync.WaitGroup
}

// Field игровое поле.
type Field struct {
	cfg          *config.Config
	appInfo      AppInfo
	metaInfo     MetaInfo
	NumberGamers int
	Objects      []systems.CommonSystem
	DeadObjects  []systems.CommonSystem
}

// New возвращает объект игрового поля.
func New(cfg *config.Config) Field {
	return Field{
		cfg: cfg,
		appInfo: AppInfo{
			mutex: &sync.Mutex{},
			jobWG: &sync.WaitGroup{},
		},
		metaInfo: MetaInfo{
			MaxNumberGamers: cfg.Game.MaxGamers,
			SizeX:           cfg.Game.SizeX,
			SizeY:           cfg.Game.SizeY,
			PreSelectPlaces: cfg.Game.PreSelectPlaces,
		},
		NumberGamers: 0,
		Objects:      nil,
		DeadObjects:  nil,
	}
}

// Info возвращает информацию об игровых объектах в свободной форме.
func (f *Field) Info() map[string]interface{} {
	result := make(map[string]interface{}, len(f.Objects))
	for _, object := range f.Objects {
		m := make(map[string]interface{})
		if obj, ok := object.(systems.PositionSystem); ok {
			m["coord"] = []int{obj.GetPosition().X, obj.GetPosition().Y}
		}

		if obj, ok := object.(systems.HealthSystem); ok {
			m["health"] = obj.GetHealth().HitPoints
		}

		obj := object.(systems.CommonSystem)

		result[obj.GetCommon().Id] = m
	}
	return result
}

// Start запускает процессы.
func (f *Field) Start(ctx context.Context) {
	for idx, obj := range f.Objects {
		if tank, ok := obj.(*entities.Tank); ok {
			tank.Position.X = f.metaInfo.PreSelectPlaces[idx][0]
			tank.Position.Y = f.metaInfo.PreSelectPlaces[idx][1]
		}
	}
	f.runJobs(ctx)
}
