package entities

import (
	"eight-stones/ecs-tank-engine/engine/components"
	"eight-stones/ecs-tank-engine/engine/config"
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

func NewTank(cfg *config.TankConfig) Tank {
	return Tank{
		Common: components.Common{
			Id: uuid.New().String(),
		},
		Position: components.Position{
			X: -1,
			Y: -1,
		},
		Movement: components.Movement{
			Recharge: &components.Recharge{
				DefaultDuration: cfg.MoveRechargeDefaultDuration,
				FreeAction:      cfg.MoveRechargeFreeAction,
				MaxAction:       cfg.MoveRechargeMaxAction,
			},
		},
		Rotatement: components.Rotatement{
			Recharge: &components.Recharge{
				DefaultDuration: cfg.RotateRechargeDefaultDuration,
				FreeAction:      cfg.RotateRechargeFreeAction,
				MaxAction:       cfg.RotateRechargeMaxAction,
			},
		},
		Health: components.Health{
			HitPoints:    cfg.HitPoints,
			MaxHitPoints: cfg.MaxHitPoints,
		},
		Damage: components.Damage{
			DamagePoints: cfg.DamagePoints,
		},
	}
}
