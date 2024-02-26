package gonimator

type TickDuration = uint16

// Part represents a single image of an animation.
type Part struct {
	// Duration in ticks
	Duration TickDuration

	// Value is the number that is returned while this part is active
	Value int
}

// NewPlayer returns a new PartPlayer based on a pointer to the part.
func (p *Part) NewPlayer() *PartPlayer {
	return &PartPlayer{
		part: p,
	}
}
