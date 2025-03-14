package engine

import (
	"context"
	"time"

	"github.com/Eight-Stones/ecs-tank-engine/v2/config"
	"github.com/Eight-Stones/ecs-tank-engine/v2/systems"
)

// Field game field.
type Field struct {
	cfg         *config.Config
	sync        syncInfo
	gameInfo    gameInfo
	cache       cache
	cancelCtx   func()
	Objects     []systems.InfoSystem
	DeadObjects []systems.InfoSystem
}

// New create new instance.
func New(opt ...config.Option) *Field {
	cfg := config.Default()
	for _, o := range opt {
		o(cfg)
	}

	field := Field{
		cfg: cfg,
		gameInfo: gameInfo{
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

func (f *Field) tryLockWithContext(ctx context.Context) bool {
	for {
		select {
		case <-ctx.Done():
			return false
		default:
			if f.sync.mutex.TryLock() {
				return true
			}
			select {
			case <-time.After(10 * time.Millisecond):
			case <-ctx.Done():
				return false
			}
		}
	}
}
