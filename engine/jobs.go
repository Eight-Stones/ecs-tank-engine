package engine

import (
	"context"
)

// runJobs start inner processes.
func (f *Field) runJobs(ctx context.Context) {
	go func() {
		f.appInfo.jobWG.Add(1)
		go f.autoReplaceDeadJob(ctx)
		f.appInfo.jobWG.Add(2)
		go f.autoMovementJob(ctx)
		f.appInfo.jobWG.Wait()
	}()
}
