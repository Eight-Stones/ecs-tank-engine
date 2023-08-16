package components

const (
	UnSuccess         = 0
	Success           = 0b1
	MoveSuccess       = 0b10
	MoveUnSuccess     = 0b100
	ShotSuccess       = 0b1000
	ShotUnSuccess     = 0b10000
	Collision         = 0b100000
	CollisionWithMove = 0b1000000
	NoneCollision     = 0b10000000
	TankNotFound      = 0b100000000
)

const (
	Undefined = iota
	Left
	Right
	Up
	Down
)

var MovementValue = map[uint][]int{
	Left:  {-1, 0},
	Right: {1, 0},
	Down:  {0, -1},
	Up:    {0, 1},
}
