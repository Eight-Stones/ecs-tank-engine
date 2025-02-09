package systems

// AutoMovementSystem describes auto moving system.
type AutoMovementSystem interface {
	CanAutoMove()
	MovementSystem
}
