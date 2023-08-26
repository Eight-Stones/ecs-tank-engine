package engine

import (
	"eight-stones/ecs-tank-engine/engine/common"
	"eight-stones/ecs-tank-engine/engine/entities"
	"eight-stones/ecs-tank-engine/engine/systems"
	"errors"
)

func (f *Field) find(id string) (systems.CommonSystem, int) {
	for idx := range f.Objects {
		if f.Objects[idx].GetCommon().Id == id {
			return f.Objects[idx], common.Found
		}
	}
	return nil, common.NotFound
}

func (f *Field) getAllCanCommon() []systems.CommonSystem {
	var result []systems.CommonSystem
	for idx := range f.Objects {
		if obj, ok := f.Objects[idx].(systems.CommonSystem); ok {
			result = append(result, obj)
		}
	}
	return result
}

func (f *Field) getAllCanPosition() []systems.PositionSystem {
	var result []systems.PositionSystem
	for idx := range f.Objects {
		if obj, ok := f.Objects[idx].(systems.PositionSystem); ok {
			result = append(result, obj)
		}
	}
	return result
}

func (f *Field) getAllCanMovement() []systems.MovementSystem {
	var result []systems.MovementSystem
	for idx := range f.Objects {
		if obj, ok := f.Objects[idx].(systems.MovementSystem); ok {
			result = append(result, obj)
		}
	}
	return result
}

func (f *Field) getAllCanAutoMovement() []systems.CommonSystem {
	var result []systems.CommonSystem
	for idx := range f.Objects {
		if obj, ok := f.Objects[idx].(systems.AutoMovementSystem); ok {
			result = append(result, obj.(systems.CommonSystem))
		}
	}
	return result
}

func (f *Field) getAllCanRecharged() []systems.CommonSystem {
	var result []systems.CommonSystem
	for idx := range f.Objects {
		if _, ok := f.Objects[idx].(systems.AutoMovementSystem); ok {
			result = append(result, f.Objects[idx])
			continue
		}
		if _, ok := f.Objects[idx].(systems.RotatementSystem); ok {
			result = append(result, f.Objects[idx])
		}
	}
	return result
}

func (f *Field) getAllCanHealth() []systems.HealthSystem {
	var result []systems.HealthSystem
	for idx := range f.Objects {
		if obj, ok := f.Objects[idx].(systems.HealthSystem); ok {
			result = append(result, obj)
		}
	}
	return result
}

func (f *Field) AddTank() (string, error) {
	if f.NumberGamers >= f.cfg.Game.MaxGamers {
		return "", errors.New("too much players")
	}
	f.NumberGamers++
	return f.addTank()
}

func (f *Field) addTank() (string, error) {
	tank := entities.NewTank(&f.cfg.Tank)
	f.Objects = append(f.Objects, &tank)
	return tank.Id, nil
}
