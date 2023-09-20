package engine

import (
	"context"
	"time"
)

// autoInformerJob start process of collection action data.
func (f *Field) autoInformerJob(ctx context.Context, out chan interface{}) {
	defer f.sync.jobWG.Done()
	ticker := time.NewTicker(time.Millisecond)
	for {
		select {
		case <-ticker.C:
			if data := f.autoInformer(ctx); data != nil {
				out <- data
			}
		case <-ctx.Done():
			return
		}
	}
}

// autoInformer retrieves data.
func (f *Field) autoInformer(_ context.Context) interface{} {
	return f.cache.read()
}
