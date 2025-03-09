package engine

import (
	"github.com/Eight-Stones/ecs-tank-engine/v2/components"
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

type Vision struct {
	Radius int
}

type Radar struct {
	Radius int
}

func (c *cache) saveCreate(
	id string,
	objectType components.ObjectType,
	direction components.Direction,
	position []int,
	health int,
) {
	create := Info{
		Id:         id,
		Type:       ActionCreate,
		ObjectType: objectType,
		MetaInfo: Create{
			Direction: direction,
			Position:  position,
			Health:    health,
		},
	}
	c.save(create)
}

func (c *cache) saveRemove(id string, objectType components.ObjectType) {
	remove := Info{
		Id:         id,
		Type:       ActionRemove,
		ObjectType: objectType,
		MetaInfo:   Remove{},
	}
	c.save(remove)
}

func (c *cache) saveRotatement(id string, objectType components.ObjectType, old, new components.Direction) {
	info := Info{
		Id:         id,
		Type:       ActionRotate,
		ObjectType: objectType,
		MetaInfo: Rotate{
			Old: old,
			New: new,
		},
	}
	c.save(info)
}

func (c *cache) saveStep(id string, objectType components.ObjectType, old, new []int) {
	info := Info{
		Id:         id,
		Type:       ActionMove,
		ObjectType: objectType,
		MetaInfo: Position{
			Old: old,
			New: new,
		},
	}
	c.save(info)
}

func (c *cache) saveCollision(id string, objectType components.ObjectType, old, new int) {
	info := Info{
		Id:         id,
		Type:       ActionHealth,
		ObjectType: objectType,
		MetaInfo: Health{
			Old: old,
			New: new,
		},
	}
	c.save(info)
}

func (c *cache) saveShoot(id string, objectType components.ObjectType, ammoLeft int, direction components.Direction) {
	info := Info{
		Id:         id,
		Type:       ActionShoot,
		ObjectType: objectType,
		MetaInfo: Shoot{
			AmmoLeft:  ammoLeft,
			Direction: direction,
		},
	}
	c.save(info)
}

func (c *cache) saveVision(id string, objectType components.ObjectType, radius int) {
	info := Info{
		Id:         id,
		Type:       ActionVision,
		ObjectType: objectType,
		MetaInfo: Vision{
			Radius: radius,
		},
	}
	c.save(info)
}

func (c *cache) saveRadar(id string, objectType components.ObjectType, radius int) {
	info := Info{
		Id:         id,
		Type:       ActionRadar,
		ObjectType: objectType,
		MetaInfo: Radar{
			Radius: radius,
		},
	}
	c.save(info)
}
