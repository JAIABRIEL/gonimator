package gonimator

type TickDuration = uint16

// Part represents a single image of an animation.
type Part[T any] struct {
	// Duration in ticks
	Duration TickDuration

	// Value is the number that is returned while this part is active
	Value T
}

// NewPlayer returns a new PartPlayer based on a pointer to the part.
func (p *Part[T]) NewPlayer() *PartPlayer[T] {
	return &PartPlayer[T]{
		part: p,
	}
}
