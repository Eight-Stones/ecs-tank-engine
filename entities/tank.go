package entities

import (
	"ecs-tank-engine/components"
	"ecs-tank-engine/config"
	"github.com/google/uuid"
)

// Tank war machine.
type Tank struct {
	components.Info
	components.Statistic
	components.Position
	components.Vision
	components.Radar
	components.Movement
	components.Rotatement
	components.Health
	components.Damage
	components.Shooting
}

// NewTank return new instance of entity.
func NewTank(cfg *config.TankConfig) Tank {
	return Tank{
		Info: components.Info{
			Id:   uuid.New().String(),
			Type: components.TypeTankId,
		},
		Position: components.Position{
			X:         -1,
			Y:         -1,
			Direction: components.Right,
		},
		Vision: components.Vision{
			Radius: cfg.Vision,
			Recharge: &components.Recharge{
				DefaultDuration: cfg.VisionRechargeDefaultDuration,
			},
		},
		Radar: components.Radar{
			Radius: cfg.Radar,
			Recharge: &components.Recharge{
				DefaultDuration: cfg.RadarRechargeDefaultDuration,
			},
		},
		Movement: components.Movement{
			Recharge: &components.Recharge{
				DefaultDuration: cfg.MoveRechargeDefaultDuration,
			},
		},
		Rotatement: components.Rotatement{
			Recharge: &components.Recharge{
				DefaultDuration: cfg.RotateRechargeDefaultDuration,
			},
		},
		Health: components.Health{
			HitPoints:    cfg.HitPoints,
			MaxHitPoints: cfg.MaxHitPoints,
		},
		Damage: components.Damage{
			DamagePoints: cfg.DamagePoints,
		},
		Shooting: components.Shooting{
			Ammo:    cfg.Ammo,
			MaxAmmo: cfg.MaxAmmo,
			Recharge: &components.Recharge{
				DefaultDuration: cfg.ShootRechargeDefaultDuration,
			},
		},
	}
}
