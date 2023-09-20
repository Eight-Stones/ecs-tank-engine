package engine

import (
	"eight-stones/ecs-tank-engine/engine/components"
)

type Create struct {
	Direction components.Direction
	Position  []int
	Health    int
}

type Remove struct{}

type Rotate struct {
	Old components.Direction
	New components.Direction
}

type Position struct {
	Old []int
	New []int
}

type Health struct {
	Old int
	New int
}

type Shoot struct {
	AmmoLeft  int
	Direction components.Direction
}

func (c *cache) saveCreate(
	id string,
	direction components.Direction,
	position []int,
	health int,
) {
	create := Info{
		Id:   id,
		Type: ActionCreate,
		MetaInfo: Create{
			Direction: direction,
			Position:  position,
			Health:    health,
		},
	}
	c.save(create)
}

func (c *cache) saveRemove(id string) {
	remove := Info{
		Id:       id,
		Type:     ActionRemove,
		MetaInfo: Remove{},
	}
	c.save(remove)
}

func (c *cache) saveRotatement(id string, old, new components.Direction) {
	info := Info{
		Id:   id,
		Type: ActionRotate,
		MetaInfo: Rotate{
			Old: old,
			New: new,
		},
	}
	c.save(info)
}

func (c *cache) saveStep(id string, old, new []int) {
	info := Info{
		Id:   id,
		Type: ActionMove,
		MetaInfo: Position{
			Old: old,
			New: new,
		},
	}
	c.save(info)
}

func (c *cache) saveCollision(id string, old, new int) {
	info := Info{
		Id:   id,
		Type: ActionHealth,
		MetaInfo: Health{
			Old: old,
			New: new,
		},
	}
	c.save(info)
}

func (c *cache) saveShoot(id string, ammoLeft int, direction components.Direction) {
	info := Info{
		Id:   id,
		Type: ActionShoot,
		MetaInfo: Shoot{
			AmmoLeft:  ammoLeft,
			Direction: direction,
		},
	}
	c.save(info)
}
