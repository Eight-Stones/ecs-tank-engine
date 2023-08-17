package engine

import (
	"eight-stones/ecs-tank-engine/engine/entities"
	"eight-stones/ecs-tank-engine/engine/systems"
	"sync"
)

type Border struct {
	X int
	Y int
}

type Params struct {
	MaxGamers int
	FieldSize int
}

type Field struct {
	mutex sync.Mutex
	Params
	Gamers  int
	Border  Border
	Objects []systems.CommonSystem
}

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

func (f *Field) Start() {
	places := [][]int{{0, 0}, {1, 0}, {15, 0}, {0, 15}} //[][]int{{0, 0}, {15, 15}, {15, 0}, {0, 15}}
	cnt := 0
	for _, obj := range f.Objects {
		if tank, ok := obj.(*entities.Tank); ok {
			tank.Position.X = places[cnt][0]
			tank.Position.Y = places[cnt][1]
			cnt++
		}
	}
}
