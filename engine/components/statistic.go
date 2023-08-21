package components

type Statistic struct {
	Actions []int
}

func (s *Statistic) GetStatistic() *Statistic { return s }
