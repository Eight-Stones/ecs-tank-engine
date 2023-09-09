package systems

import "eight-stones/ecs-tank-engine/engine/components"

// InfoSystem describes info system.
type InfoSystem interface {
	GetInfo() *components.Info
}
