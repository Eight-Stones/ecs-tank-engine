package entities

import "eight-stones/ecs-tank-engine/engine/components"

type Wall struct {
	components.Info
	components.Position
}

type Tank struct {
	components.Info
	components.Position
	components.Movement
}

type Bullet struct {
	components.Info
	components.Position
	components.Movement
	components.AutoMovement
}
