package systems

import "eight-stones/ecs-tank-engine/engine/components"

type PositionSystem interface {
	GetPosition() *components.Position
}

func ChangePosition(incX, incY int, in PositionSystem) {
	position := in.GetPosition()
	position.X += incX
	position.Y += incY
}
