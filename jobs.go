package engine

import (
	"context"
)

// runJobs start inner processes.
func (f *Field) runJobs(ctx context.Context) {
	go func() {
		f.sync.jobWG.Add(1)
		go f.autoReplaceDeadJob(ctx)
		f.sync.jobWG.Add(1)
		go f.autoMovementJob(ctx)
		f.sync.jobWG.Add(1)
		go f.autoInformerJob(ctx, f.cache.getOut())
	}()
}
