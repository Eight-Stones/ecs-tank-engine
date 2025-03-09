package engine

import (
	"context"
	"time"

	"github.com/Eight-Stones/ecs-tank-engine/v2/common"
	"github.com/Eight-Stones/ecs-tank-engine/v2/pkg/utils"
)

// autoMovementJob  start process of changing placement ob object which can automove.
func (f *Field) autoMovementJob(ctx context.Context) {
	defer f.sync.jobWG.Done()
	ticker := time.NewTicker(f.cfg.Game.Jobs.AutoMover)
	for {
		select {
		case <-ticker.C:
			f.autoMovement(ctx)
		case <-ctx.Done():
			return
		}
	}
}

// autoMovementJob change placement ob object which can automove.
func (f *Field) autoMovement(_ context.Context) {
	f.sync.mutex.Lock()
	defer f.sync.mutex.Unlock()
	now := time.Now()
	objects := f.getAllCanAutoMovement()
	for i := range objects {
		code := f.move(objects[i], now)
		if utils.CheckBitMask(code, common.OkCollision) || utils.CheckBitMask(code, common.NotInterruptOkCollision) {
			f.replaceDeadById(objects[i].GetInfo().Id)
		}
	}
}
