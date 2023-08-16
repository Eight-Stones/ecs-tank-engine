package common

const (
	UndefinedDoing           = 0
	BreakBorder              = 0b1
	NoneBreakBorder          = 0b10
	MoveSuccess              = 0b100
	MoveUnSuccess            = 0b1000
	NoneCollision            = 0b10000
	CollisionSuccess         = 0b100000
	CollisionSuccessWithMove = 0b1000000
	ShotSuccess              = 0b10000000
	ShotUnSuccess            = 0b1000000000
	TankFound                = 0b10000000000
	TankNotFound             = 0b100000000000
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
