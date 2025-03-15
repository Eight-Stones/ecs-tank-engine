package config

import (
	"time"
)

type Option func(*Config)

// WithGameMaxGamers sets value of maximum player that can be added in game.
func WithGameMaxGamers(in int) Option {
	return func(o *Config) {
		o.Game.MaxGamers = in
	}
}

// WithGameFieldSizeX sets value x-axis of game field.
func WithGameFieldSizeX(in int) Option {
	return func(o *Config) {
		o.Game.SizeX = in
	}
}

// WithGameFieldSizeY sets value y-axis of game field.
func WithGameFieldSizeY(in int) Option {
	return func(o *Config) {
		o.Game.SizeY = in
	}
}

// WithGamePreSelectPlaces sets start places for players.
func WithGamePreSelectPlaces(in [][]int) Option {
	return func(o *Config) {
		o.Game.PreSelectPlaces = in
	}
}

// WithGamePreSelectDirections sets start directions for players.
func WithGamePreSelectDirections(in []Direction) Option {
	return func(o *Config) {
		o.Game.PreSelectDirection = in
	}
}

// WithJobsAutoMoverDuration sets ticker duration for auto move job.
func WithJobsAutoMoverDuration(in time.Duration) Option {
	return func(o *Config) {
		o.Game.Jobs.AutoMover = in
	}
}

// WithJobsRechargerDuration set recharge waiting job duration.
//
// Now its not affected on game node.
func WithJobsRechargerDuration(in time.Duration) Option {
	return func(o *Config) {
		o.Game.Jobs.Recharger = in
	}
}

// WithJobsReplacerDuration sets duration for replace objects job.
func WithJobsReplacerDuration(in time.Duration) Option {
	return func(o *Config) {
		o.Game.Jobs.Replacer = in
	}
}

// WithTankStartHP sets value of start hit points of player's tank.
func WithTankStartHP(in int) Option {
	return func(o *Config) {
		o.Tank.HitPoints = in
	}
}

// WithTankMaxHP sets maximum hit points for player's tank.
func WithTankMaxHP(in int) Option {
	return func(o *Config) {
		o.Tank.MaxHitPoints = in
	}
}

// WithTankMoveRecharge sets how often player can move his tank.
func WithTankMoveRecharge(in time.Duration) Option {
	return func(o *Config) {
		o.Tank.MoveRechargeDefaultDuration = in
	}
}

// WithTankRotateRecharge sets how often player can rotate his tank.
func WithTankRotateRecharge(in time.Duration) Option {
	return func(o *Config) {
		o.Tank.RotateRechargeDefaultDuration = in
	}
}

// WithTankShootRecharge sets how ofte player can shoot by his tank.
func WithTankShootRecharge(in time.Duration) Option {
	return func(o *Config) {
		o.Tank.ShootRechargeDefaultDuration = in
	}
}

// WithTankDamage set level of damage from collision of player's tank.
func WithTankDamage(in int) Option {
	return func(o *Config) {
		o.Tank.DamagePoints = in
	}
}

// WithTankAmmo sets start value ammo for player's tank.
func WithTankAmmo(in int) Option {
	return func(o *Config) {
		o.Tank.Ammo = in
	}
}

// WithTankMaxAmmo sets maximum value ammo for player's tank.
func WithTankMaxAmmo(in int) Option {
	return func(o *Config) {
		o.Tank.MaxAmmo = in
	}
}

// WithTankVision sets radius of vision for player's tank.
func WithTankVision(in int) Option {
	return func(o *Config) {
		o.Tank.Vision = in
	}
}

// WithTankVisionRecharge sets recharge of vision for player's tank.
func WithTankVisionRecharge(in time.Duration) Option {
	return func(o *Config) {
		o.Tank.VisionRechargeDefaultDuration = in
	}
}

// WithTankRadar sets radius of radar for player's tank.
func WithTankRadar(in int) Option {
	return func(o *Config) {
		o.Tank.Radar = in
	}
}

// WithTankRadarRecharge sets recharge of radar for player's tank.
func WithTankRadarRecharge(in time.Duration) Option {
	return func(o *Config) {
		o.Tank.RadarRechargeDefaultDuration = in
	}
}

// WithBulletHP sets bullet hit points.
func WithBulletHP(in int) Option {
	return func(o *Config) {
		o.Bullet.HitPoints = in
	}
}

// WithBulletMaxHP sets bullet maximum hit points.
func WithBulletMaxHP(in int) Option {
	return func(o *Config) {
		o.Bullet.MaxHitPoints = in
	}
}

// WithBulletMoveRecharge sets recharge of bullet move.
func WithBulletMoveRecharge(in time.Duration) Option {
	return func(o *Config) {
		o.Bullet.MoveRechargeDefaultDuration = in
	}
}

// WithBulletDamage sets damage of bullet.
func WithBulletDamage(in int) Option {
	return func(o *Config) {
		o.Bullet.DamagePoints = in
	}
}
