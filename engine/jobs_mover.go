package engine

import (
	"context"
	"eight-stones/ecs-tank-engine/engine/common"
	"eight-stones/ecs-tank-engine/engine/pkg/utils"
	"time"
)

// autoMovementJob  start process of changing placement ob object which can automove.
func (f *Field) autoMovementJob(ctx context.Context) {
	defer f.appInfo.jobWG.Done()
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
	f.appInfo.mutex.Lock()
	defer f.appInfo.mutex.Unlock()
	now := time.Now()
	objects := f.getAllCanAutoMovement()
	for i := range objects {
		code := f.move(objects[i], now)
		if utils.CheckBitMask(code, common.OkCollision) || utils.CheckBitMask(code, common.NotInterruptOkCollision) {
			f.replaceDeadById(objects[i].GetInfo().Id)
		}
	}
}
