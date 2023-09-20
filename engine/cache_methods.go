package engine

import "eight-stones/ecs-tank-engine/engine/common"

type Create struct {
	Direction common.Direction
	Position  []int
	Health    int
}

type Remove struct{}

type Rotate struct {
	Old common.Direction
	New common.Direction
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
	Position  []int
	Direction uint
}

func (c *cache) saveCreate(
	id string,
	direction common.Direction,
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
		Type:     ActionCreate,
		MetaInfo: Remove{},
	}
	c.save(remove)
}

func (c *cache) saveRotatement(id string, old, new common.Direction) {
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
