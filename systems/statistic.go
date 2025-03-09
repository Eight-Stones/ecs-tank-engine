package systems

import "github.com/Eight-Stones/ecs-tank-engine/v2/components"

// StatisticSystem decribes statistic system.
type StatisticSystem interface {
	GetStatistic() *components.Statistic
}

// AddAction added action code in list of actions.
func AddAction(code int, in InfoSystem) {
	if in.GetInfo().Parent != nil {
		// TODO добавить сюда дополнительный код, который будет говорить о дочерности объекта
		AddAction(code, in.GetInfo().Parent)
		return
	}
	statistic, ok := in.(StatisticSystem)
	if !ok {
		return
	}
	statistic.GetStatistic().Actions = append(statistic.GetStatistic().Actions, code)
}
