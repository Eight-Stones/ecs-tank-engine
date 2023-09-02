package engine

import (
	"eight-stones/ecs-tank-engine/engine/common"
	"eight-stones/ecs-tank-engine/engine/entities"
	"eight-stones/ecs-tank-engine/engine/systems"
)

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

	if obj, ok := object.(systems.DamageSystem); ok {
		state[common.KeyStatDamage] = obj.GetDamage().DamagePoints
	}

	return state
}

// CurrentState return game current state.
func (f *Field) CurrentState() map[string]map[string]interface{} {
	result := make(map[string]map[string]interface{}, len(f.Objects))
	for _, object := range f.Objects {
		state := getState(object)
		obj := object.(systems.CommonSystem)
		result[obj.GetCommon().Id] = state
	}
	return result
}

func (f *Field) Result() map[string]map[string]interface{} {
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
