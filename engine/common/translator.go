package common

import (
	"fmt"
	"strings"
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
	NoCollision:             "отсутстие столкновения",
	NotInterruptOkCollision: "наличие столкновение без прерывания",
	OkShot:                  "выстрел успешен",
	FailShot:                "выстрел неуспешен",
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
	NoCollision,
	NotInterruptOkCollision,
	OkStep,
	FailStep,
	OkShot,
	FailShot,
	Damaged,
	BothDamaged,
	Fail,
	Ok,
	Ban,
}

func TranslatePrint(actions int) string {
	var result string
	for idx := range order {
		if actions&order[idx] == order[idx] {
			result += fmt.Sprintf("%v->", aliases[order[idx]])
		}
	}
	return result
}

func TranslateBuilder(actions int, builder *strings.Builder) {
	for idx := range order {
		if actions&order[idx] == order[idx] {
			builder.WriteString(fmt.Sprintf("%v->", aliases[order[idx]]))
		}
	}
}
