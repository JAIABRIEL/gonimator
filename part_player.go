package gonimator

// PartPlayer counts to a parts duration.
// After the duration is reached it just resets itself.
type PartPlayer struct {
	part    *Part
	current uint16
}

// Update increases the parts current state by one tick.
// Returns false if the end is reached.
// It's a loop so just update it again to restart.
func (p *PartPlayer) Update() bool {
	p.current = (p.current + 1) % p.part.Duration
	return p.current != 0
}

// Cancel sets the current state to zero.
// Doesn't need to be called after Update returned zero.
func (p *PartPlayer) Cancel() {
	p.current = 0
}

// Get returns the parts value.
func (p *PartPlayer) Get() int {
	return p.part.Value
}
