package components

type Common struct {
	Id string
}

func (c *Common) GetCommon() *Common { return c }
