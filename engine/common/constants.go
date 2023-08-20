package common

const (
	UndefinedDoing     = 0
	Fail               = 0b1
	Success            = 0b10
	Border             = 0b100
	Step               = 0b1000
	Rotate             = 0b10000
	Collision          = 0b100000
	CollisionSuccess   = 0b1000000
	NotInterruptAction = 0b100000000
	Damaged            = 0b1000000000
	BothDamaged        = 0b10000000000
	Shot               = 0b1000000000000
	NotFound           = 0b10000000000000
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
