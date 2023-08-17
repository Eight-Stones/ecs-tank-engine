package entities

import "eight-stones/ecs-tank-engine/engine/components"

type Wall struct {
	components.Common
	components.Position
}

type Tank struct {
	components.Common
	components.Position
	components.Movement
	components.Rotatement
	components.Health
	components.Damage
}

type Bullet struct {
	components.Common
	components.Position
	components.Movement
	components.AutoMovement
	components.Health
	components.Damage
}
