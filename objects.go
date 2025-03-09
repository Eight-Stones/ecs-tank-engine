package engine

import (
	"errors"

	"github.com/Eight-Stones/ecs-tank-engine/v2/common"
	"github.com/Eight-Stones/ecs-tank-engine/v2/entities"
	"github.com/Eight-Stones/ecs-tank-engine/v2/systems"
)

// find finds object by his id.
func (f *Field) find(id string) (systems.InfoSystem, int) {
	for idx := range f.Objects {
		if f.Objects[idx].GetInfo().Id == id {
			return f.Objects[idx], common.Found
		}
	}
	return nil, common.NotFound
}

// getAllCanCommon get all common object,
func (f *Field) getAllCanCommon() []systems.InfoSystem {
	var result []systems.InfoSystem
	for idx := range f.Objects {
		if obj, ok := f.Objects[idx].(systems.InfoSystem); ok {
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
func (f *Field) getAllCanAutoMovement() []systems.InfoSystem {
	var result []systems.InfoSystem
	for idx := range f.Objects {
		if obj, ok := f.Objects[idx].(systems.AutoMovementSystem); ok {
			result = append(result, obj.(systems.InfoSystem))
		}
	}
	return result
}

// getAllCanRecharged get all recharged object.
func (f *Field) getAllCanRecharged() []systems.InfoSystem {
	var result []systems.InfoSystem
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
	if f.gameInfo.NumberGamers >= f.cfg.Game.MaxGamers {
		return "", errors.New("too much players")
	}
	f.gameInfo.NumberGamers++
	return f.addTank()
}

// addTank added new tank to game.
func (f *Field) addTank() (string, error) {
	tank := entities.NewTank(&f.cfg.Tank)
	f.Objects = append(f.Objects, &tank)
	return tank.Id, nil
}
