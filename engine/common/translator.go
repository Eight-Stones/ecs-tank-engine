package common

import "fmt"

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
	Ban:                     "запрет действия",
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
	Ban,
}

func Translate(actions int) string {
	var result string
	for idx := range order {
		if actions&order[idx] == order[idx] {
			result += fmt.Sprintf("%v->", aliases[order[idx]])
		}
	}
	return result
}
