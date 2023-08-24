package engine

import (
	"context"
)

// runJobs запускает все побочные процессы игры.
func (f *Field) runJobs(ctx context.Context) {
	go func() {
		f.appInfo.jobWG.Add(1)
		go f.autoReplaceDeadJob(ctx)
		f.appInfo.jobWG.Wait()
	}()
}

// autoMovementJob изменяет координаты объектов, которые могут двигаться самостоятельно.
func (f *Field) autoMovementJob(ctx context.Context) {

}
