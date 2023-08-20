package common

const (
	UndefinedDoing                   = 0
	BreakBorder                      = 0b1
	NoneBreakBorder                  = 0b10
	MoveSuccess                      = 0b100
	MoveUnSuccess                    = 0b1000
	RotateSuccess                    = 0b10000
	RotateUnSuccess                  = 0b100000
	NoneCollision                    = 0b1000000
	CollisionSuccess                 = 0b10000000
	CollisionSuccessNotInterruptMove = 0b100000000
	Damaged                          = 0b1000000000
	BothDamaged                      = 0b10000000000
	ShotSuccess                      = 0b1000000000000
	ShotUnSuccess                    = 0b10000000000000
	TankFound                        = 0b100000000000000
	TankNotFound                     = 0b1000000000000000
)

const (
	Undefined = iota
	Left
	Right
	Up
	Down
)

const (
	DoNothing           = 0
	CanOnlyDamage       = 0b1
	CanOnlyDamaged      = 0b10
	CanDamagedAndDamage = 0b11
)

const (
	CollisionBothDamage   = 0b1
	CollisionFirstDamage  = 0b10
	CollisionSecondDamage = 0b100
)

var MovementValue = map[uint][]int{
	Left:  {-1, 0},
	Right: {1, 0},
	Down:  {0, -1},
	Up:    {0, 1},
}
