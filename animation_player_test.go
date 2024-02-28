package gonimator_test

import (
	"math/rand"
	"testing"

	"github.com/JAIABRIEL/gonimator"
)

func TestAnimationPlayer(t *testing.T) {
	partAmount := rand.Intn(100) + 1
	a := &gonimator.Animation[int]{
		Parts: make([]*gonimator.Part[int], partAmount),
	}

	totalDuration := 0

	for i := range a.Parts {
		duration := rand.Intn(100) + 1
		value := rand.Intn(10000)
		totalDuration += duration

		a.Parts[i] = NewPart(uint16(duration), value)
	}

	ap := a.NewPlayer()

	total := 0

	for ap.Update() {
		total++
	}

	if !ap.Update() {
		t.Error("AnimationPlayer reset failed")
	}

	// use total + 1 since the last ap.Update() returns false but still updates
	if totalDuration != (total + 1) {
		t.Error("Mismatch in expected and result duration")
	}
}
