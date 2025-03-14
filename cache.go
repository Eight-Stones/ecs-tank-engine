package engine

import (
	"sync"

	"github.com/Eight-Stones/ecs-tank-engine/v2/components"
)

type ActionType uint

func (at ActionType) String() string {
	switch at {
	case ActionCreate:
		return "create"
	case ActionRemove:
		return "remove"
	case ActionRotate:
		return "rotate"
	case ActionMove:
		return "move"
	case ActionShoot:
		return "shoot"
	case ActionHealth:
		return "health"
	case ActionVision:
		return "vision"
	case ActionRadar:
		return "radar"
	}
	return ""
}

const (
	ActionCreate ActionType = iota
	ActionRemove
	ActionRotate
	ActionMove
	ActionShoot
	ActionHealth
	ActionVision
	ActionRadar
)

// Info main log about actions.
type Info struct {
	Id         string
	Type       ActionType
	ObjectType components.ObjectType
	MetaInfo   interface{}
}

// cache stores all detail actions.
type cache struct {
	mu   sync.Mutex
	out  chan Info
	idx  int
	data []Info
}

// init initialize cache.
func (c *cache) init() {
	c.mu = sync.Mutex{}
	c.out = make(chan Info, 100)
	c.idx = 0
	c.data = make([]Info, 0, 100)
}

// save saves data on cache.
func (c *cache) save(in Info) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data = append(c.data, in)
}

// read return data element from cache.
func (c *cache) read() *Info {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.idx >= len(c.data) {
		return nil
	}
	defer func() { c.idx++ }()
	return &c.data[c.idx]
}

// getOut return chan with detail action data.
func (c *cache) getOut() chan Info {
	return c.out
}

func (c *cache) close() {
	close(c.out)
}
