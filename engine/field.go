package engine

import (
	"eight-stones/ecs-tank-engine/engine/config"
	"eight-stones/ecs-tank-engine/engine/systems"
)

// Field game field.
type Field struct {
	cfg         *config.Config
	sync        syncInfo
	gameInfo    GameInfo
	cache       Cache
	Objects     []systems.InfoSystem
	DeadObjects []systems.InfoSystem
}

// New create new instance.
func New(cfg *config.Config) *Field {
	field := Field{
		cfg: cfg,
		gameInfo: GameInfo{
			MaxNumberGamers:     cfg.Game.MaxGamers,
			SizeX:               cfg.Game.SizeX,
			SizeY:               cfg.Game.SizeY,
			PreSelectPlaces:     cfg.Game.PreSelectPlaces,
			PreSelectDirections: cfg.Game.PreSelectDirection,
		},
	}
	field.sync.init()
	field.cache.init()
	return &field
}
