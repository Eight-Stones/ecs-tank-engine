package engine

import (
	"context"
	"eight-stones/ecs-tank-engine/engine/common"
	"eight-stones/ecs-tank-engine/engine/entities"
	"eight-stones/ecs-tank-engine/engine/pkg/helper"
	"eight-stones/ecs-tank-engine/engine/systems"
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
			helper.DrawField(f.metaInfo.SizeX, f.metaInfo.SizeY, f.CurrentState())
		}
	}
}

func (f *Field) DrawResult() {
	helper.DrawResult(f.ResultState())
}

func getState(object systems.CommonSystem) map[string]interface{} {
	state := make(map[string]interface{})

	switch object.(type) {
	case *entities.Tank:
		state[common.KeyStatObjectKind] = common.KeyObjectKindTank
	case *entities.Bullet:
		state[common.KeyStatObjectKind] = common.KeyObjectKindBullet
	}

	if obj, ok := object.(systems.PositionSystem); ok {
		state[common.KeyStatPositionCoordinate] = []int{obj.GetPosition().X, obj.GetPosition().Y}
	}

	if obj, ok := object.(systems.MovementSystem); ok {
		state[common.KeyStatMovementDirection] = obj.GetPosition().Direction
	}

	if obj, ok := object.(systems.HealthSystem); ok {
		state[common.KeyStatHitPoints] = obj.GetHealth().HitPoints
	}

	return state
}

// CurrentState returns game current state.
func (f *Field) CurrentState() map[string]map[string]interface{} {
	result := make(map[string]map[string]interface{}, len(f.Objects))
	for _, object := range f.Objects {
		state := getState(object)
		obj := object.(systems.CommonSystem)
		result[obj.GetCommon().Id] = state
	}
	return result
}

// ResultState returns game result state.
func (f *Field) ResultState() map[string]map[string]interface{} {
	result := make(map[string]map[string]interface{}, len(f.Objects))
	for _, object := range append(f.Objects, f.DeadObjects...) {
		tank, ok := object.(*entities.Tank)
		if !ok {
			continue
		}
		state := getState(tank)
		state[common.KeyStatActions] = tank.GetStatistic().Actions
		result[tank.GetCommon().Id] = state
	}

	return result
}
