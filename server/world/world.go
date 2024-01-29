package world

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/aquilax/go-perlin"
	"github.com/dynamitemc/dynamite/server/world/anvil"
	"github.com/dynamitemc/dynamite/server/world/generator"
)

type World struct {
	//name of the folder which holds the contents of the world
	name string

	overWorld *Dimension
}

func OpenWorld(name string) (*World, error) {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		w := &World{
			name:      name,
			overWorld: NewDimension("minecraft:overworld"),
		}

		w.overWorld.rd = anvil.NewReader(name + "/region/")
		w.overWorld.gen = &generator.Default{Perlin: perlin.NewPerlin(2, 2, 1, rand.Int63())}
		fmt.Println("created world in memory")
		return w, nil
	}

	return nil, nil
}

func (w *World) OverWorld() *Dimension {
	return w.overWorld
}
