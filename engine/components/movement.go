package components

type Movement struct {
	Direction uint
}

func (m *Movement) GetMovement() *Movement { return m }

type AutoMovement struct{}

func (am *AutoMovement) CanAutoMove() {}
