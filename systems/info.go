package systems

import "ecs-tank-engine/components"

// InfoSystem describes info system.
type InfoSystem interface {
	GetInfo() *components.Info
}
