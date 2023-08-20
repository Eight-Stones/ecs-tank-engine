package components

type Common struct {
	Id     string
	Parent *Common
}

func (c *Common) GetCommon() *Common { return c }

type Statistic struct{}

func (s *Statistic) GetStatistic() *Statistic { return s }
