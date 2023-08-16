package engine

import (
	"eight-stones/ecs-tank-engine/engine/common"
	"eight-stones/ecs-tank-engine/engine/components"
	"eight-stones/ecs-tank-engine/engine/entities"
	"eight-stones/ecs-tank-engine/engine/systems"
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"sync"
)

type Border struct {
	X int
	Y int
}

type Params struct {
	MaxGamers int
	FieldSize int
}

type Field struct {
	mutex sync.Mutex
	Params
	Gamers  int
	Border  Border
	Objects []systems.PositionSystem
}

func (f *Field) find(id string) (*entities.Tank, int) {
	for idx := range f.Objects {
		obj, ok := f.Objects[idx].(*entities.Tank)
		if !ok {
			continue
		}
		if obj.Info.Id == id {
			return obj, common.TankFound
		}
	}
	return nil, common.TankNotFound
}

func (f *Field) getAllCanPosition() []systems.PositionSystem {
	var result []systems.PositionSystem
	for idx := range f.Objects {
		if obj, ok := f.Objects[idx].(systems.PositionSystem); ok {
			result = append(result, obj)
		}
	}
	return result
}

func (f *Field) getAllCanMovement() []systems.MovementSystem {
	var result []systems.MovementSystem
	for idx := range f.Objects {
		if obj, ok := f.Objects[idx].(systems.MovementSystem); ok {
			result = append(result, obj)
		}
	}
	return result
}

func (f *Field) getAllCanAutoMovement() []systems.AutoMovementSystem {
	var result []systems.AutoMovementSystem
	for idx := range f.Objects {
		if obj, ok := f.Objects[idx].(systems.AutoMovementSystem); ok {
			result = append(result, obj)
		}
	}
	return result
}

func (f *Field) AddTank() (string, error) {
	myuuid, err := uuid.NewV4()
	if err != nil {
		return "", fmt.Errorf("uuid.NewV4:%w", err)
	}

	if f.Gamers >= f.MaxGamers {
		return "", errors.New("too much players")
	}

	f.Objects = append(f.Objects, &entities.Tank{
		Info: components.Info{
			Id: myuuid.String(),
		},
		Position: components.Position{
			X: -1,
			Y: -1,
		},
		Movement: components.Movement{
			Direction: common.Undefined,
		},
	})

	f.Gamers++

	return myuuid.String(), nil

}

func (f *Field) Info() map[string][]int {
	result := make(map[string][]int, len(f.Objects))
	for _, obj := range f.Objects {
		result["t"] = []int{obj.GetPosition().X, obj.GetPosition().Y}
	}
	return result
}

func (f *Field) Start() {
	places := [][]int{{0, 0}, {15, 15}, {15, 0}, {0, 15}}
	cnt := 0
	for _, obj := range f.Objects {
		if tank, ok := obj.(*entities.Tank); ok {
			tank.Position.X = places[cnt][0]
			tank.Position.Y = places[cnt][1]
			cnt++
		}
	}
}

func (f *Field) checkBorder(direction uint, in systems.PositionSystem) int {
	incX, incY := systems.GetIncrementMoveSystem(direction)
	object := in.GetPosition()
	if object.X+incX < 0 || object.X+incX > f.Border.X || object.Y+incY < 0 || object.Y+incY > f.Border.Y {
		return common.BreakBorder
	}
	return common.NoneBreakBorder
}

func (f *Field) checkCollision(direction uint, moved, fixed systems.PositionSystem) int {
	return systems.CheckCollision(direction, moved, fixed)
}

func (f *Field) checkAutoCollision() {}

func (f *Field) Move(id string, direction uint) int {
	tank, code := f.find(id)
	doing := 0b0 | code
	if code&common.TankNotFound == common.TankNotFound {
		return doing
	}

	doing = doing | f.checkBorder(direction, tank)
	if doing&common.BreakBorder == common.BreakBorder {
		return doing
	}

	canPositionObjects := f.getAllCanPosition()
	for _, obj := range canPositionObjects {
		code = systems.CheckCollision(direction, tank, obj)
		switch code {
		case systems.Success:
			doing = doing | common.CollisionSuccess
			// TODO post collision logic
		case systems.Fail:
		}
	}

	if doing&common.CollisionSuccessWithMove == 0 && doing&common.CollisionSuccess == 0 {
		doing = doing | common.NoneCollision
	}

	systems.StepMoveSystem(direction, tank)

	doing = doing | common.MoveSuccess

	return doing
}
