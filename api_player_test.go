package engine

import (
	"sync"
	"testing"
	"time"

	"ecs-tank-engine/common"
	"ecs-tank-engine/components"
	"ecs-tank-engine/entities"
	"ecs-tank-engine/systems"
)

func TestField_Rotate(t *testing.T) {
	id := "test_tank"
	tank := &entities.Tank{
		Info: components.Info{
			Id: id,
		},
		Statistic: components.Statistic{},
		Position: components.Position{
			Direction: components.Down,
		},
		Rotatement: components.Rotatement{
			Recharge: &components.Recharge{
				Until:           time.Now().Add(time.Second * 30),
				DefaultDuration: time.Millisecond,
			},
		},
		Health: components.Health{
			HitPoints:    1,
			MaxHitPoints: 1,
		},
	}

	f := &Field{
		cfg: nil,
		sync: syncInfo{
			mutex: &sync.Mutex{},
			jobWG: &sync.WaitGroup{},
		},
		gameInfo: gameInfo{
			NumberGamers:        1,
			MaxNumberGamers:     1,
			SizeX:               15,
			SizeY:               15,
			PreSelectPlaces:     nil,
			PreSelectDirections: nil,
		},
		cache:     cache{},
		cancelCtx: nil,
		Objects: []systems.InfoSystem{
			tank,
		},
		DeadObjects: nil,
	}

	code := f.Rotate("not_found", components.Up)
	if code&common.NotFound != common.NotFound {
		t.Errorf("component was found by id='%v'", "not_found")
	}

	code = f.Rotate(id, components.Up)
	if code&common.Ban != common.Ban {
		t.Errorf("component was found by id='%v'", "not_found")
	}

	for _, direction := range []components.Direction{components.Up, components.Right, components.Down, components.Left} {
		tank.Rotatement.Until = time.Now().Add(time.Second * -1) // reset recharge
		code := f.Rotate(id, direction)
		if code&common.Found != common.Found {
			t.Errorf("component was found by id='%v'", id)
		}
		if code&common.OkRotate != common.OkRotate {
			t.Errorf("component id='%v' cant rotate to direction %v", id, direction)
		}

	}
}
