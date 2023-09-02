package engine

import (
	"eight-stones/ecs-tank-engine/engine/common"
	"eight-stones/ecs-tank-engine/engine/entities"
	"eight-stones/ecs-tank-engine/engine/systems"
	"errors"
)

// find finds object by his id.
func (f *Field) find(id string) (systems.CommonSystem, int) {
	for idx := range f.Objects {
		if f.Objects[idx].GetCommon().Id == id {
			return f.Objects[idx], common.Found
		}
	}
	return nil, common.NotFound
}

// getAllCanCommon get all common object,
func (f *Field) getAllCanCommon() []systems.CommonSystem {
	var result []systems.CommonSystem
	for idx := range f.Objects {
		if obj, ok := f.Objects[idx].(systems.CommonSystem); ok {
			result = append(result, obj)
		}
	}
	return result
}

// getAllCanPosition get all position object.
func (f *Field) getAllCanPosition() []systems.PositionSystem {
	var result []systems.PositionSystem
	for idx := range f.Objects {
		if obj, ok := f.Objects[idx].(systems.PositionSystem); ok {
			result = append(result, obj)
		}
	}
	return result
}

// getAllCanAutoMovement get all automovement object.
func (f *Field) getAllCanAutoMovement() []systems.CommonSystem {
	var result []systems.CommonSystem
	for idx := range f.Objects {
		if obj, ok := f.Objects[idx].(systems.AutoMovementSystem); ok {
			result = append(result, obj.(systems.CommonSystem))
		}
	}
	return result
}

// getAllCanRecharged get all recharged object.
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

// AddTank added new tank to game.
func (f *Field) AddTank() (string, error) {
	if f.NumberGamers >= f.cfg.Game.MaxGamers {
		return "", errors.New("too much players")
	}
	f.NumberGamers++
	return f.addTank()
}

// addTank added new tank to game.
func (f *Field) addTank() (string, error) {
	tank := entities.NewTank(&f.cfg.Tank)
	f.Objects = append(f.Objects, &tank)
	return tank.Id, nil
}
