package common

import "eight-stones/ecs-tank-engine/engine/components"

const (
	Fail                    = 0b1
	Ok                      = 0b10
	Ban                     = 0b100
	OkBorder                = 0b1000
	FailBorder              = 0b10000
	OkStep                  = 0b100000
	FailStep                = 0b1000000
	OkRotate                = 0b10000000
	FailRotate              = 0b100000000
	OkCollision             = 0b1000000000
	NoCollision             = 0b10000000000
	NotInterruptOkCollision = 0b1000000000000
	Damaged                 = 0b10000000000000
	BothDamaged             = 0b100000000000000
	OkShot                  = 0b10000000000000000
	FailShot                = 0b100000000000000000
	Disappear               = 0b1000000000000000000
	OkVision                = 0b10000000000000000000
	FailVision              = 0b100000000000000000000
	OkRadar                 = 0b1000000000000000000000
	FailRadar               = 0b10000000000000000000000
	Found                   = 0b100000000000000000000000
	NotFound                = 0b1000000000000000000000000
)

const (
	DoNothing           = 0
	CanOnlyDamage       = 0b1
	CanOnlyDamaged      = 0b10
	CanDamagedAndDamage = 0b11
)

var MovementValue = map[components.Direction][]int{
	components.Left:  {-1, 0},
	components.Right: {1, 0},
	components.Down:  {0, -1},
	components.Up:    {0, 1},
}
