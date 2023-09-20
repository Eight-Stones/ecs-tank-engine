package engine

import (
	"eight-stones/ecs-tank-engine/engine/common"
	"eight-stones/ecs-tank-engine/engine/entities"
	"eight-stones/ecs-tank-engine/engine/systems"
)

// defineType defines type of entity.
func defineType(object systems.InfoSystem) string {
	switch object.(type) {
	case *entities.Tank:
		return common.KeyObjectKindTank
	case *entities.Bullet:
		return common.KeyObjectKindBullet
	}
	return ""
}

// getState collect and return full info about entity.
func getState(object systems.InfoSystem) map[string]interface{} {
	state := make(map[string]interface{})

	state[common.KeyStatObjectKind] = defineType(object)

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
		obj := object.(systems.InfoSystem)
		result[obj.GetInfo().Id] = state
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
		result[tank.GetInfo().Id] = state
	}

	return result
}
