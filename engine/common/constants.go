package common

const (
	Fail                    = 0b1
	Ok                      = 0b10
	OkBorder                = 0b100
	FailBorder              = 0b1000
	OkStep                  = 0b10000
	FailStep                = 0b100000
	OkRotate                = 0b1000000
	FailRotate              = 0b10000000
	OkCollision             = 0b100000000
	FailCollision           = 0b1000000000
	NotInterruptOkCollision = 0b100000000000
	Damaged                 = 0b1000000000000
	BothDamaged             = 0b10000000000000
	OkShot                  = 0b1000000000000000
	FailShot                = 0b10000000000000000
	Found                   = 0b100000000000000000
	NotFound                = 0b1000000000000000000
)

var aliases = map[int]string{
	Fail:                    "неудача",
	Ok:                      "успех",
	OkBorder:                "отсутствие нарушения границ карты",
	FailBorder:              "попытка выйти за границы",
	OkStep:                  "успешное движение",
	FailStep:                "неудачное движение",
	OkRotate:                "успешный поворот",
	FailRotate:              "неудачный поворот",
	OkCollision:             "наличие столкновление",
	FailCollision:           "отсутстие столкновения",
	NotInterruptOkCollision: "наличие столкновение без прерывания",
	Damaged:                 "получение урона",
	BothDamaged:             "совместное получение урона",
	Found:                   "объект найден",
	NotFound:                "объект не найден",
}

var order = []int{
	Found,
	NotFound,
	OkRotate,
	FailRotate,
	OkBorder,
	FailBorder,
	OkCollision,
	FailCollision,
	NotInterruptOkCollision,
	OkStep,
	FailStep,
	Damaged,
	BothDamaged,
	Fail,
	Ok,
}

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
