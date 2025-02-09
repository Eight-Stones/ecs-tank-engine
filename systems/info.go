package systems

import "github.com/Eight-Stones/ecs-tank-engine/components"

// InfoSystem describes info system.
type InfoSystem interface {
	GetInfo() *components.Info
}
