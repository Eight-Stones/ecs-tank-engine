package components

import "time"

type Recharge struct {
	Until           time.Time
	DefaultDuration time.Duration
	FreeAction      int
	MaxAction       int
}

func (r *Recharge) GetRecharge() *Recharge { return r }

func (r *Recharge) IncFreeAction() {
	current := r.FreeAction + 1
	if current > r.MaxAction {
		return
	}
	r.FreeAction = current
}

func (r *Recharge) DecFreeAction() {
	current := r.FreeAction - 1
	if current < 0 {
		return
	}
	r.FreeAction = current
}

func (r *Recharge) SetUntil(now time.Time) {
	r.Until = now.Add(r.DefaultDuration)
}
