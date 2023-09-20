package components

import "time"

// Recharge adds the ability for actions to recharge and be performed a certain number of times.
type Recharge struct {
	Until           time.Time
	DefaultDuration time.Duration
}

// GetRecharge returns Recharge component.
func (r *Recharge) GetRecharge() *Recharge { return r }

// IsRechargeDone checks that recharging is done.
func (r *Recharge) IsRechargeDone(now time.Time) bool {
	return r.Until.Before(now)
}

// SetUntil set end recharging time.
func (r *Recharge) SetUntil(now time.Time) {
	r.Until = now.Add(r.DefaultDuration)
}
