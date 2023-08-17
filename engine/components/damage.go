package components

import "time"

type Damage struct {
	DamagePoints int
}

func (d *Damage) GetDamage() *Damage { return d }

type Shooting struct {
	Ammo    int
	MaxAmmo int
	Refill  time.Time
}

func (s *Shooting) GetShooting() *Shooting { return s }
func (s *Shooting) HasAmmo() bool          { return s.Ammo > 0 }
