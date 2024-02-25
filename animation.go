package gonimator

// Animation sticks a collection of parts together.
type Animation struct {
	// Parts that will be played
	Parts []*Part
}

// NewPlayer returns a new AnimationPlayer based on a pointer to the animation.
func (a *Animation) NewPlayer() *AnimationPlayer {
	ap := &AnimationPlayer{
		parts: make([]*PartPlayer, len(a.Parts)),
	}

	for i := range ap.parts {
		ap.parts[i] = a.Parts[i].NewPlayer()
	}

	return ap
}
