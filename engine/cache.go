package engine

import (
	"sync"
)

type ActionType uint

const (
	ActionRotate ActionType = iota
	ActionMove
	ActionShoot
	ActionHealth
)

type Info struct {
	Id       string
	Type     uint
	MetaInfo interface{}
}

type Rotate struct {
	Old uint
	New uint
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

// Cache stores all detail actions.
type Cache struct {
	mu   sync.Mutex
	out  chan interface{}
	idx  int
	data []interface{}
}

// init initialize cache.
func (c *Cache) init() {
	c.mu = sync.Mutex{}
	c.out = make(chan interface{})
	c.idx = 0
	c.data = make([]interface{}, 100)
}

// save saves data on cache.
func (c *Cache) save(in interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data = append(c.data, in)
}

// read return data element from cache.
func (c *Cache) read() interface{} {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.idx >= len(c.data) {
		return nil
	}
	defer func() { c.idx++ }()
	return c.data[c.idx]
}

// getOut return chan with detail action data.
func (c *Cache) getOut() chan interface{} {
	return c.out
}
