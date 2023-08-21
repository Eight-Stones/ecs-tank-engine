package components

import "time"

// Recharge компонент, добавляющий возможность действиям перезаряжаться и выполняться определенное количество раз.
type Recharge struct {
	Until           time.Time
	DefaultDuration time.Duration
}

// GetRecharge интерфейс-маркер компонента.
func (r *Recharge) GetRecharge() *Recharge { return r }

// IsRechargeDone проверяет, что действие перезарядилось.
func (r *Recharge) IsRechargeDone(now time.Time) bool {
	return r.Until.Before(now)
}

// SetUntil устанавливает время перезарядки time.Now() + DefaultDuration.
func (r *Recharge) SetUntil(now time.Time) {
	r.Until = now.Add(r.DefaultDuration)
}
