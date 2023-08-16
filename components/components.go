package components

type Position struct {
	X int
	Y int
}

func (p *Position) GetPosition() *Position { return p }

type Movement struct {
	Direction uint
	AutoMove  bool
}

func (m *Movement) CanMovementByItself() bool { return m.AutoMove }
func (m *Movement) GetMovement() *Movement    { return m }

type Border struct {
	X int
	Y int
}

type PositionSystem interface {
	GetPosition() *Position
}

func GetIncrementMoveSystem(direction uint) (x, y int) {
	values := MovementValue[direction]
	return values[0], values[1]
}

func CheckBorder(direction uint, b *Border, in PositionSystem) int {
	incX, incY := GetIncrementMoveSystem(direction)
	object := in.GetPosition()
	if object.X+incX < 0 || object.X+incX > b.X || object.Y+incY < 0 || object.Y+incY > b.Y {
		return UnSuccess
	}
	return Success
}

func CheckCollision(direction uint, inMove, inCllsn PositionSystem) int {
	incX, incY := GetIncrementMoveSystem(direction)
	pMove := inMove.GetPosition()
	pCllsn := inCllsn.GetPosition()
	if pCllsn.X == pMove.X+incX && pCllsn.Y == pMove.Y+incY {
		if false {
			return CollisionWithMove
		}
		return Collision
	}
	return NoneCollision
}

type MovementSystem interface {
	GetMovement() *Movement
	PositionSystem
}

type AutoMovementSystem interface {
	CanMovementByItself() bool
	MovementSystem
}

func ChangePosition(incX, incY int, in PositionSystem) {
	position := in.GetPosition()
	position.X += incX
	position.Y += incY
}

func StepMoveSystem(direction uint, in MovementSystem) {
	incX, incY := GetIncrementMoveSystem(direction)
	in.GetMovement().Direction = direction
	ChangePosition(incX, incY, in.GetPosition())
}
