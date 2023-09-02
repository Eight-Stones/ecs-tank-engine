package engine

import (
	"context"
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
		f.move(objects[i], now)
	}
}
