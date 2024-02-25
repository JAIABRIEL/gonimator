package gonimator_test

import (
	"math/rand"
	"testing"

	"github.com/JAIABRIEL/gonimator"
)

func TestPartPlayer(t *testing.T) {
	duration := rand.Intn(5000) + 1
	value := rand.Intn(100000)
	pp := NewPart(uint16(duration), value).NewPlayer()

	i := 0
	for i < duration {
		i++

		r := pp.Update()
		if r == false && i != duration {
			t.Error("Mismatch between part and tested duration")
		}

		if pp.Get() != value {
			t.Error("Got wrong value")
		}
	}
}

func NewPart(duration uint16, value int) *gonimator.Part {
	return &gonimator.Part{
		Duration: duration,
		Value:    value,
	}
}
