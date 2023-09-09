package components

// Movement describes the ability of an entity to move.
type Movement struct {
	*Recharge
}

// GetMovement returns Movement component.
func (m *Movement) GetMovement() *Movement { return m }

// Rotatement describes the ability of an entity to change direction..
type Rotatement struct {
	*Recharge
}

// GetRotatement returns Rotatement component.
func (r *Rotatement) GetRotatement() *Rotatement { return r }

// NotInterruptMovement adds the ability to not interrupt movement in the event of a collision.
type NotInterruptMovement struct{}

// GetNotInterruptMovement returns NotInterruptMovement component.
func (nim *NotInterruptMovement) GetNotInterruptMovement() {}

// AutoMovement describes the ability to move independently in the direction of an entity.
type AutoMovement struct{}

// CanAutoMove returns AutoMovement component.
func (am *AutoMovement) CanAutoMove() {}
