package components

import (
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"sync"
)

type Info struct {
	Id string
}

func (i *Info) GetInfo() *Info { return i }

type Tank struct {
	Info
	Position
	Movement
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
	Objects []PositionSystem
}

func (f *Field) find(id string) *Tank {
	for idx := range f.Objects {
		obj, ok := f.Objects[idx].(*Tank)
		if !ok {
			continue
		}
		if obj.Info.Id == id {
			return obj
		}
	}
	return nil
}

func (f *Field) getAllCanPosition() []PositionSystem {
	var result []PositionSystem
	for idx := range f.Objects {
		if obj, ok := f.Objects[idx].(PositionSystem); ok {
			result = append(result, obj)
		}
	}
	return result
}

func (f *Field) getAllCanMovement() []MovementSystem {
	var result []MovementSystem
	for idx := range f.Objects {
		if obj, ok := f.Objects[idx].(MovementSystem); ok {
			result = append(result, obj)
		}
	}
	return result
}

func (f *Field) getAllCanAutoMovement() []AutoMovementSystem {
	var result []AutoMovementSystem
	for idx := range f.Objects {
		if obj, ok := f.Objects[idx].(AutoMovementSystem); ok {
			if !obj.CanMovementByItself() {
				continue
			}
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

	f.Objects = append(f.Objects, &Tank{
		Info: Info{
			Id: myuuid.String(),
		},
		Position: Position{
			X: -1,
			Y: -1,
		},
		Movement: Movement{
			Direction: Undefined,
			AutoMove:  false,
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
		if tank, ok := obj.(*Tank); ok {
			tank.Position.X = places[cnt][0]
			tank.Position.Y = places[cnt][1]
			cnt++
		}
	}
}

func (f *Field) Move(id string, direction uint) int {
	tank := f.find(id)
	if tank == nil {
		return TankNotFound
	}

	doing := 0b0
	if code := CheckBorder(direction, &f.Border, tank); code == UnSuccess {
		return MoveUnSuccess
	}

	canPositionObjects := f.getAllCanPosition()
	for _, obj := range canPositionObjects {
		code := CheckCollision(direction, tank, obj)
		doing = doing | code
		switch code {
		case Collision, CollisionWithMove:
			doing = doing | Collision
			// TODO add damage logic and return NEW code
			break
		case NoneCollision:
			continue
		}
	}

	if doing&CollisionWithMove == 0 && doing&Collision == 0 {
		doing = doing | NoneCollision
	}

	StepMoveSystem(direction, tank)

	doing = doing | MoveSuccess

	return doing
}
