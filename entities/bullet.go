package entities

import (
	"time"

	"github.com/Eight-Stones/ecs-tank-engine/v2/components"
	"github.com/Eight-Stones/ecs-tank-engine/v2/config"
	"github.com/Eight-Stones/ecs-tank-engine/v2/systems"
	"github.com/google/uuid"
)

// Bullet tank shell.
type Bullet struct {
	components.Info
	components.Position
	components.Movement
	components.AutoMovement
	components.NotInterruptMovement
	components.Health
	components.Damage
}

// NewBullet return new instance of entity.
func NewBullet(
	cfg *config.BulletConfig,
	parent systems.InfoSystem,
	x, y int,
	direction components.Direction,
) Bullet {
	return Bullet{
		Info: components.Info{
			Id:     uuid.New().String(),
			Type:   components.TypeBulletId,
			Parent: parent.(*components.Info),
		},
		Position: components.Position{
			X:         x,
			Y:         y,
			Direction: direction,
		},
		Movement: components.Movement{
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
