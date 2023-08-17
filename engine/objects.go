package engine

import (
	"eight-stones/ecs-tank-engine/engine/common"
	"eight-stones/ecs-tank-engine/engine/components"
	"eight-stones/ecs-tank-engine/engine/entities"
	"eight-stones/ecs-tank-engine/engine/systems"
	"errors"
	"github.com/google/uuid"
)

func (f *Field) find(id string) (*entities.Tank, int) {
	for idx := range f.Objects {
		obj, ok := f.Objects[idx].(*entities.Tank)
		if !ok {
			continue
		}
		if obj.Common.Id == id {
			return obj, common.TankFound
		}
	}
	return nil, common.TankNotFound
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
	if f.Gamers >= f.MaxGamers {
		return "", errors.New("too much players")
	}

	tank := &entities.Tank{
		Common: components.Common{
			Id: uuid.New().String(),
		},
		Position: components.Position{
			X: -1,
			Y: -1,
		},
		Movement: components.Movement{
			Direction: common.Undefined,
		},
		Health: components.Health{
			HitPoints:    100,
			MaxHitPoints: 150,
		},
		Damage: components.Damage{
			DamagePoints: 30,
		},
	}

	f.Objects = append(f.Objects, tank)

	f.Gamers++

	return tank.Id, nil
}
