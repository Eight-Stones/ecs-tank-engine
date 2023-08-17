package components

type Movement struct {
	Direction uint
}

func (m *Movement) GetMovement() *Movement { return m }

type Rotatement struct{}

func (r *Rotatement) CanRotate() {}

type NotInterruptMovement struct{}

func (nim *NotInterruptMovement) DoesNotInterruptMovement() {}

type AutoMovement struct{}

func (am *AutoMovement) CanAutoMove() {}
