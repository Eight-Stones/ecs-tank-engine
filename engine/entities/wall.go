package entities

import "eight-stones/ecs-tank-engine/engine/components"

type Wall struct {
	components.Common
	components.Position
}
