package engine

import (
	"context"
	"eight-stones/ecs-tank-engine/engine/common"
	"eight-stones/ecs-tank-engine/engine/config"
	"eight-stones/ecs-tank-engine/engine/entities"
	"eight-stones/ecs-tank-engine/engine/pkg/helper"
	"eight-stones/ecs-tank-engine/engine/systems"
	"sync"
	"time"
)

// MetaInfo метаинформация.
type MetaInfo struct {
	NumberGamers        int
	MaxNumberGamers     int
	SizeX               int
	SizeY               int
	PreSelectPlaces     [][]int
	PreSelectDirections []uint
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

// Info возвращает информацию об игровых объектах в свободной форме.
func (f *Field) Info() map[string]map[string]interface{} {
	result := make(map[string]map[string]interface{}, len(f.Objects))
	for _, object := range f.Objects {
		m := make(map[string]interface{})

		switch object.(type) {
		case *entities.Tank:
			m[common.KeyObjectKind] = common.KeyObjectTank
		case *entities.Bullet:
			m[common.KeyObjectKind] = common.KeyObjectBullet
		}

		if obj, ok := object.(systems.PositionSystem); ok {
			m[common.KeyPositionCoordinate] = []int{obj.GetPosition().X, obj.GetPosition().Y}
		}

		if obj, ok := object.(systems.MovementSystem); ok {
			m[common.KeyMovementDirection] = obj.GetMovement().Direction
		}

		if obj, ok := object.(systems.HealthSystem); ok {
			m[common.KeyStatHitPoints] = obj.GetHealth().HitPoints
		}

		if obj, ok := object.(systems.DamageSystem); ok {
			m[common.KeyStatDamage] = obj.GetDamage().DamagePoints
		}

		obj := object.(systems.CommonSystem)

		result[obj.GetCommon().Id] = m
	}
	return result
}

func (f *Field) DrawConsole(ctx context.Context) {
	ticker := time.NewTicker(time.Millisecond * 100)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			helper.DrawField(f.metaInfo.SizeX, f.metaInfo.SizeY, f.Info())
		}
	}
}

// Start запускает процессы.
func (f *Field) Start(ctx context.Context) {
	for idx, obj := range f.Objects {
		if tank, ok := obj.(*entities.Tank); ok {
			tank.Position.X = f.metaInfo.PreSelectPlaces[idx][0]
			tank.Position.Y = f.metaInfo.PreSelectPlaces[idx][1]
			tank.Movement.Direction = f.metaInfo.PreSelectDirections[idx]
		}
	}
	f.runJobs(ctx)
}
