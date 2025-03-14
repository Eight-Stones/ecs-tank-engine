package engine

import (
	"context"
	"time"
)

// autoInformerJob start process of collection action data.
func (f *Field) autoInformerJob(ctx context.Context, out chan Info) {
	defer f.sync.jobWG.Done()
	ticker := time.NewTicker(time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if data := f.autoInformer(ctx); data != nil {
				out <- *data // TODO make all as pointer
			}
		case <-ctx.Done():
			return
		}
	}
}

// autoInformer retrieves data.
func (f *Field) autoInformer(_ context.Context) *Info {
	return f.cache.read()
}
