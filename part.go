package ganimator

// Part represents a single image of an animation.
type Part struct {
	// Duration in ticks
	Duration uint16

	// Value is the number that is returned while this part is active
	Value int
}
