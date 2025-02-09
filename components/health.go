package components

// Health describes the health level and adds the ability to take damage.
type Health struct {
	HitPoints    int
	MaxHitPoints int
}

// GetHealth returns Health component.
func (h *Health) GetHealth() *Health { return h }
