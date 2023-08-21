package config

import "time"

type Config struct {
	Game GameConfig
	Tank TankConfig
}

type JobsConfig struct {
	AutoMover time.Duration
	Recharger time.Duration
	Replacer  time.Duration
}

type GameConfig struct {
	MaxGamers       int
	SizeX           int
	SizeY           int
	PreSelectPlaces [][]int
	Jobs            JobsConfig
}

type TankConfig struct {
	HitPoints                     int
	MaxHitPoints                  int
	MoveRechargeDefaultDuration   time.Duration
	RotateRechargeDefaultDuration time.Duration
	DamagePoints                  int
}

func Default() *Config {
	return &Config{
		Game: GameConfig{
			MaxGamers:       4,
			SizeX:           15,
			SizeY:           15,
			PreSelectPlaces: [][]int{{0, 0}, {15, 15}, {15, 0}, {0, 15}},
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
			DamagePoints:                  30,
		},
	}
}
