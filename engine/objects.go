package engine

import (
	"eight-stones/ecs-tank-engine/engine/common"
	"eight-stones/ecs-tank-engine/engine/entities"
	"eight-stones/ecs-tank-engine/engine/systems"
	"errors"
)

func (f *Field) find(id string) (*entities.Tank, int) {
	for idx := range f.Objects {
		obj, ok := f.Objects[idx].(*entities.Tank)
		if !ok {
			continue
		}
		if obj.Common.Id == id {
			return obj, common.Found
		}
	}
	return nil, common.NotFound
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

func (f *Field) getAllCanAutoMovement() []systems.AutoMovementSystem {
	var result []systems.AutoMovementSystem
	for idx := range f.Objects {
		if obj, ok := f.Objects[idx].(systems.AutoMovementSystem); ok {
			result = append(result, obj)
		}
	}
	return result
}

func (f *Field) AddTank() (string, error) {
	if f.NumberGamers >= f.cfg.Game.MaxGamers {
		return "", errors.New("too much players")
	}

	tank := entities.NewTank(&f.cfg.Tank)
	f.Objects = append(f.Objects, &tank)
	f.NumberGamers++

	return tank.Id, nil
}
