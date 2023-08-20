package components

type Movement struct {
	Direction uint
	Recharge
}

func (m *Movement) GetMovement() *Movement { return m }

type Rotatement struct {
	Recharge
}

func (r *Rotatement) GetRotatement() *Rotatement { return r }

type NotInterruptMovement struct{}

func (nim *NotInterruptMovement) DoesNotInterruptMovement() {}

type AutoMovement struct{}

func (am *AutoMovement) CanAutoMove() {}
