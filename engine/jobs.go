package engine

import (
	"context"
)

// runJobs запускает все побочные процессы игры.
func (f *Field) runJobs(ctx context.Context) {
	go func() {
		//f.appInfo.jobWG.Add(1)
		//f.appInfo.jobWG.Wait()
	}()
}

// autoMovementJob изменяет координаты объектов, которые могут двигаться самостоятельно.
func (f *Field) autoMovementJob(ctx context.Context) {

}

// autoReplaceDeadJob переносит мертвые игровые объекты в лог объектов.
func (f *Field) autoReplaceDeadJob(ctx context.Context) {

}
