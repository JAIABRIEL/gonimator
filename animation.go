package gonimator

// Animation sticks a collection of parts together.
type Animation[T any] struct {
	// Parts that will be played
	Parts []*Part[T]
}

// NewPlayer returns a new AnimationPlayer based on a pointer to the animation.
func (a *Animation[T]) NewPlayer() *AnimationPlayer[T] {
	ap := &AnimationPlayer[T]{
		parts: make([]*PartPlayer[T], len(a.Parts)),
	}

	for i := range ap.parts {
		ap.parts[i] = a.Parts[i].NewPlayer()
	}

	return ap
}
