package entities

import (
	"eight-stones/ecs-tank-engine/engine/components"
	"eight-stones/ecs-tank-engine/engine/config"
	"eight-stones/ecs-tank-engine/engine/systems"
	"github.com/google/uuid"
	"time"
)

type Bullet struct {
	components.Common
	components.Position
	components.Movement
	components.AutoMovement
	components.Health
	components.Damage
}

func NewBullet(cfg *config.BulletConfig, parent systems.CommonSystem, x, y int, direction uint) Bullet {
	return Bullet{
		Common: components.Common{
			Id:     uuid.New().String(),
			Parent: parent.(*components.Common),
		},
		Position: components.Position{
			X: x,
			Y: y,
		},
		Movement: components.Movement{
			Direction: direction,
			Recharge: &components.Recharge{
				Until:           time.Time{},
				DefaultDuration: cfg.MoveRechargeDefaultDuration,
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
