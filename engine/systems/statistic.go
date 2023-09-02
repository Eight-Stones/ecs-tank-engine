package systems

import "eight-stones/ecs-tank-engine/engine/components"

type StatisticSystem interface {
	GetStatistic() *components.Statistic
}

func AddActionStatisticSystem(code int, in CommonSystem) {
	if in.GetCommon().Parent != nil {
		AddActionStatisticSystem(code, in.GetCommon().Parent)
		return
	}
	statistic, ok := in.(StatisticSystem)
	if !ok {
		return
	}
	statistic.GetStatistic().Actions = append(statistic.GetStatistic().Actions, code)
}
