package config

import "time"

type Config struct {
	Game GameConfig
	Tank TankConfig
}

type GameConfig struct {
	MaxGamers       int
	SizeX           int
	SizeY           int
	PreSelectPlaces [][]int
}

type TankConfig struct {
	HitPoints                     int
	MaxHitPoints                  int
	MoveRechargeDefaultDuration   time.Duration
	MoveRechargeFreeAction        int
	MoveRechargeMaxAction         int
	RotateRechargeDefaultDuration time.Duration
	RotateRechargeFreeAction      int
	RotateRechargeMaxAction       int
	DamagePoints                  int
}

func Default() *Config {
	return &Config{
		Game: GameConfig{
			MaxGamers:       4,
			SizeX:           15,
			SizeY:           15,
			PreSelectPlaces: [][]int{{0, 0}, {15, 15}, {15, 0}, {0, 15}},
		},
		Tank: TankConfig{
			HitPoints:                     100,
			MaxHitPoints:                  150,
			MoveRechargeDefaultDuration:   time.Second,
			MoveRechargeFreeAction:        0,
			MoveRechargeMaxAction:         1,
			RotateRechargeDefaultDuration: time.Millisecond * 500,
			RotateRechargeFreeAction:      0,
			RotateRechargeMaxAction:       4,
			DamagePoints:                  30,
		},
	}
}
