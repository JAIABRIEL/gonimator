package gonimator

// PartPlayer counts to a parts duration.
// After the duration is reached it just resets itself.
type PartPlayer[T any] struct {
	part    *Part[T]
	current uint16
}

// Update increases the parts current state by one tick.
// Returns false if the end is reached.
// It's a loop so just update it again to restart.
func (p *PartPlayer[T]) Update() bool {
	p.current = (p.current + 1) % p.part.Duration
	return p.current != 0
}

// Cancel sets the current state to zero.
// Doesn't need to be called after Update returned zero.
func (p *PartPlayer[T]) Cancel() {
	p.current = 0
}

// Get returns the parts value.
func (p *PartPlayer[T]) Get() T {
	return p.part.Value
}
