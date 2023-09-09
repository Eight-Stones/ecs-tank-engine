package config

import (
	"eight-stones/ecs-tank-engine/engine/common"
	"time"
)

// Config main config.
type Config struct {
	Game   GameConfig
	Tank   TankConfig
	Bullet BulletConfig
}

// JobsConfig describes jobs parameters.
type JobsConfig struct {
	AutoMover time.Duration
	Recharger time.Duration
	Replacer  time.Duration
}

// GameConfig game parameters.
type GameConfig struct {
	MaxGamers          int
	SizeX              int
	SizeY              int
	PreSelectPlaces    [][]int
	PreSelectDirection []uint
	Jobs               JobsConfig
}

// TankConfig tank-entity params.
type TankConfig struct {
	HitPoints                     int
	MaxHitPoints                  int
	MoveRechargeDefaultDuration   time.Duration
	RotateRechargeDefaultDuration time.Duration
	ShootRechargeDefaultDuration  time.Duration
	DamagePoints                  int
	Ammo                          int
	MaxAmmo                       int
}

// BulletConfig bullet-entity params.
type BulletConfig struct {
	HitPoints                   int
	MaxHitPoints                int
	MoveRechargeDefaultDuration time.Duration
	DamagePoints                int
}

// Default return default value of config.
func Default() *Config {
	return &Config{
		Game: GameConfig{
			MaxGamers:          4,
			SizeX:              15,
			SizeY:              15,
			PreSelectPlaces:    [][]int{{0, 0}, {14, 0}, {14, 0}, {0, 14}},
			PreSelectDirection: []uint{common.Right, common.Left, common.Left, common.Right},
			Jobs: JobsConfig{
				AutoMover: time.Second,
				Recharger: time.Millisecond * 1,
				Replacer:  time.Millisecond * 1,
			},
		},
		Tank: TankConfig{
			HitPoints:                     100,
			MaxHitPoints:                  150,
			MoveRechargeDefaultDuration:   time.Second,
			RotateRechargeDefaultDuration: time.Millisecond * 500,
			ShootRechargeDefaultDuration:  time.Second,
			DamagePoints:                  20,
			Ammo:                          20,
			MaxAmmo:                       30,
		},
		Bullet: BulletConfig{
			HitPoints:                   1,
			MaxHitPoints:                1,
			MoveRechargeDefaultDuration: time.Millisecond * 300,
			DamagePoints:                35,
		},
	}
}
