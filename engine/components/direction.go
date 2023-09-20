package components

type Direction uint

func (d Direction) String() string {
	switch d {
	case Left:
		return "left"
	case Right:
		return "right"
	case Up:
		return "up"
	case Down:
		return "down"
	}
	return ""
}

const (
	Undefined Direction = iota
	Left
	Right
	Up
	Down
)
