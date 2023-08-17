package components

type Health struct {
	HitPoints    int
	MaxHitPoints int
}

func (h *Health) GetHealth() *Health { return h }
