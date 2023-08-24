package config

import (
	"eight-stones/ecs-tank-engine/engine/common"
	"time"
)

type Config struct {
	Game   GameConfig
	Tank   TankConfig
	Bullet BulletConfig
}

type JobsConfig struct {
	AutoMover time.Duration
	Recharger time.Duration
	Replacer  time.Duration
}

type GameConfig struct {
	MaxGamers          int
	SizeX              int
	SizeY              int
	PreSelectPlaces    [][]int
	PreSelectDirection []uint
	Jobs               JobsConfig
}

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

type BulletConfig struct {
	HitPoints                   int
	MaxHitPoints                int
	MoveRechargeDefaultDuration time.Duration
	DamagePoints                int
}

func Default() *Config {
	return &Config{
		Game: GameConfig{
			MaxGamers:          4,
			SizeX:              15,
			SizeY:              15,
			PreSelectPlaces:    [][]int{{0, 0}, {14, 14}, {14, 0}, {0, 14}},
			PreSelectDirection: []uint{common.Right, common.Left, common.Left, common.Right},
			Jobs: JobsConfig{
				AutoMover: time.Second,
				Recharger: time.Millisecond * 100,
				Replacer:  time.Millisecond * 50,
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
			MoveRechargeDefaultDuration: time.Millisecond * 500,
			DamagePoints:                35,
		},
	}
}
