package components

type Position struct {
	X         int
	Y         int
	Direction uint
}

func (p *Position) GetPosition() *Position { return p }
