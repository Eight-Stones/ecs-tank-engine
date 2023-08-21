package systems

import "eight-stones/ecs-tank-engine/engine/components"

type StatisticSystem interface {
	GetStatistic() *components.Statistic
}
