package entities

import "eight-stones/ecs-tank-engine/engine/components"

type Bullet struct {
	components.Common
	components.Position
	components.Movement
	components.AutoMovement
	components.Health
	components.Damage
}
