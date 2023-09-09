package systems

import "eight-stones/ecs-tank-engine/engine/components"

// StatisticSystem decribes statistic system.
type StatisticSystem interface {
	GetStatistic() *components.Statistic
}

// AddAction added action code in list of actions.
func AddAction(code int, in InfoSystem) {
	if in.GetInfo().Parent != nil {
		AddAction(code, in.GetInfo().Parent)
		return
	}
	statistic, ok := in.(StatisticSystem)
	if !ok {
		return
	}
	statistic.GetStatistic().Actions = append(statistic.GetStatistic().Actions, code)
}
