package setup

import (
	"github.com/EzyBakeOven/Prophecy/entities/humaniods"
)

// BuildWorld : Builds the world for the simulation.
func BuildWorld() {
	
	// Create one thousand entities
	for index := 0; index < 1000; index++ {
		human := human{name: "bob", age: 32}
	}
}
