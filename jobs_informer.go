package engine

import (
	"context"
	"time"
)

func (f *Field) autoInformerJob(ctx context.Context, out chan Info) {
	defer f.sync.jobWG.Done()
	ticker := time.NewTicker(time.Millisecond * 100) // Увеличиваем интервал тикера
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if data := f.autoInformer(ctx); data != nil {
				select {
				case out <- *data:
				case <-ctx.Done():
					return
				}
			}
		case <-ctx.Done():
			return
		}
	}
}

func (f *Field) autoInformer(ctx context.Context) *Info {
	select {
	case <-ctx.Done():
		return nil // Прерываем, если контекст отменен
	default:
		return f.cache.read()
	}
}
