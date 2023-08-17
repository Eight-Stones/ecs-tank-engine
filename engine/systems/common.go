package systems

import "eight-stones/ecs-tank-engine/engine/components"

type CommonSystem interface {
	GetCommon() *components.Common
}
