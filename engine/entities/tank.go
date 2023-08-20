package entities

import (
	"eight-stones/ecs-tank-engine/engine/components"
	"github.com/google/uuid"
)

type Tank struct {
	components.Common
	components.Position
	components.Movement
	components.Rotatement
	components.Health
	components.Damage
}

func NewTank() Tank {
	return Tank{
		Common: components.Common{
			Id:     uuid.New().String(),
			Parent: nil,
		},
		Position: components.Position{
			X: -1,
			Y: -1,
		},
		Movement: components.Movement{
			Direction: 0,
			Recharge: components.Recharge{
				DefaultDuration: 0,
				FreeAction:      0,
				MaxAction:       0,
			},
		},
		Rotatement: components.Rotatement{
			Recharge: components.Recharge{
				DefaultDuration: 0,
				FreeAction:      0,
				MaxAction:       0,
			},
		},
		Health: components.Health{
			HitPoints:    0,
			MaxHitPoints: 0,
		},
		Damage: components.Damage{
			DamagePoints: 0,
		},
	}
}
