package engine

import (
	"context"
	"eight-stones/ecs-tank-engine/engine/pkg/helper"
	"time"
)

// DrawConsole helper method for drawing game.
func (f *Field) DrawConsole(ctx context.Context) {
	ticker := time.NewTicker(time.Millisecond * 100)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			helper.DrawField(f.gameInfo.SizeX, f.gameInfo.SizeY, f.CurrentState())
		}
	}
}

func (f *Field) DrawResult() {
	helper.DrawResult(f.ResultState())
}
