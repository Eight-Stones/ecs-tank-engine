package components

// Statistic stores action statistics
type Statistic struct {
	Actions []int
}

// GetStatistic returns Statistic component.
func (s *Statistic) GetStatistic() *Statistic { return s }
