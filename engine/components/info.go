package components

type ObjectType uint

const (
	TypeTank ObjectType = iota
	TypeBullet
)

// Info the main component of any entity&
type Info struct {
	Id     string
	Type   ObjectType
	Parent *Info
}

// GetInfo returns Info component.
func (c *Info) GetInfo() *Info { return c }
