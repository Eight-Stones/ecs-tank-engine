package engine

import (
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

// Field игровое поле.
type Field struct {
	mutex        sync.Mutex
	cfg          *config.Config
	inf          MetaInfo
	NumberGamers int
	Objects      []systems.CommonSystem
}

// New возвращает объект игрового поля.
func New(cfg *config.Config) Field {
	return Field{
		mutex:        sync.Mutex{},
		cfg:          cfg,
		NumberGamers: 0,
		inf: MetaInfo{
			MaxNumberGamers: cfg.Game.MaxGamers,
			SizeX:           cfg.Game.SizeX,
			SizeY:           cfg.Game.SizeY,
			PreSelectPlaces: cfg.Game.PreSelectPlaces,
		},
		Objects: nil,
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
func (f *Field) Start() {
	for idx, obj := range f.Objects {
		if tank, ok := obj.(*entities.Tank); ok {
			tank.Position.X = f.inf.PreSelectPlaces[idx][0]
			tank.Position.Y = f.inf.PreSelectPlaces[idx][1]
		}
	}
	// TODO start jobs
}

// Stop останавливает процессы.
func (f *Field) Stop() {

}
