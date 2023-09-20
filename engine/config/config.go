package config

import (
	"eight-stones/ecs-tank-engine/engine/components"
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
	PreSelectDirection []components.Direction
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
	Vision                        int
	VisionRechargeDefaultDuration time.Duration
	Radar                         int
	RadarRechargeDefaultDuration  time.Duration
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
			PreSelectPlaces:    [][]int{{0, 0}, {14, 0}, {14, 14}, {0, 14}},
			PreSelectDirection: []components.Direction{components.Right, components.Left, components.Left, components.Right},
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
			Vision:                        3,
			VisionRechargeDefaultDuration: time.Millisecond * 500,
			Radar:                         6,
			RadarRechargeDefaultDuration:  time.Second * 2,
		},
		Bullet: BulletConfig{
			HitPoints:                   1,
			MaxHitPoints:                1,
			MoveRechargeDefaultDuration: time.Millisecond * 150,
			DamagePoints:                35,
		},
	}
}
