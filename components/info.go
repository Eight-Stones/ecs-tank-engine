package components

type ObjectType uint

func (ot ObjectType) String() string {
	switch ot {
	case TypeTankId:
		return TypeTank
	case TypeBulletId:
		return TypeBullet
	}
	return ""
}

const (
	UndefinedType ObjectType = iota
	TypeTankId
	TypeBulletId
)

const (
	TypeTank   = "tank"
	TypeBullet = "bullet"
)

// Info the main component of any entity&
type Info struct {
	Id     string
	Type   ObjectType
	Parent *Info
}

// GetInfo returns Info component.
func (c *Info) GetInfo() *Info { return c }
