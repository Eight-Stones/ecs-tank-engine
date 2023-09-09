package components

// Info the main component of any entity&
type Info struct {
	Id     string
	Parent *Info
}

// GetInfo returns Info component.
func (c *Info) GetInfo() *Info { return c }
