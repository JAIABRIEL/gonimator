package gonimator

// AnimationPlayer plays each of its parts for the parts duration.
type AnimationPlayer struct {
	parts []*PartPlayer

	current int
}

// Update increases the current parts state by one tick.
// Returns false if all parts reached their end.
// It's a loop so just update it again to restart.
func (ap *AnimationPlayer) Update() bool {
	if ap.parts[ap.current].Update() {
		return true
	}
	ap.current = (ap.current + 1) % len(ap.parts)

	return ap.current != 0
}

// Get returns the current parts value.
func (ap *AnimationPlayer) Get() int {
	return ap.parts[ap.current].Get()
}

// Cancel resets the animation.
func (ap *AnimationPlayer) Cancel() {
	// if used correctly, every part that is not the current one
	// should be at 0.
	ap.parts[ap.current].Cancel()
	ap.current = 0
}
