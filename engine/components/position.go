package components

type Position struct {
	X int
	Y int
}

func (p *Position) GetPosition() *Position { return p }
