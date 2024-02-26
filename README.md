# gonimator

Simple animation util library written in Go.
Define animations in your project and store them for example in your
resource library. 
I wrote this to create animations in [ebiten](https://github.com/hajimehoshi/ebiten),
based on a index of my TSX-files.

Tested on hundreds of entities with a single `AnimationPlayer` for each of them with more than 1200 TPS so
it should be pretty solid.


## Will be implemented soon

* Implement a manager to switch between different animations or define for example `Idle` animations
* Use generics for Parts.Value


## Usage

```go
package main

import "github.com/JAIABRIEL/gonimator"

// Define your animations
var myAnimation1 *gonimator.Animation = &gonimator.Animation{
	Parts: []*gonimator.Part{
		{Value: 1, Duration: 20},
		{Value: 2, Duration: 20},
		{Value: 3, Duration: 20},
		{Value: 4, Duration: 20},
		{Value: 5, Duration: 20},
		{Value: 6, Duration: 20},
	},
}

func main() {
	ap := myAnimation1.NewPlayer()
    
    // Call ap.Update() with every tick of your game.
	for ap.Update() {
		// Use value to display the image
		// Could be defined as index for tsx files.
		ap.Get()
	}
	// Animation done.
	// It's a loop so using ap.Update() again will reset everything.
}
```
