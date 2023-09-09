package components

// Damage describes the addition of the ability to deal damage.
type Damage struct {
	DamagePoints int
}

// GetDamage return Damage component.
func (d *Damage) GetDamage() *Damage { return d }

// Shooting describes the addition of the ability to shoot and store ammo.
type Shooting struct {
	Ammo    int
	MaxAmmo int
	Recharge
}

// GetShooting return Shooting component.
func (s *Shooting) GetShooting() *Shooting { return s }

// HasAmmo checks that component has ammunition.
func (s *Shooting) HasAmmo() bool { return s.Ammo > 0 }
