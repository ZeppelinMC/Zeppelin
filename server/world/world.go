package world

import "github.com/zeppelinmc/zeppelin/server/world/region"

type World struct {
	Level
	Overworld *region.Dimension

	path string
}

func NewWorld(path string) (*World, error) {
	var err error
	w := &World{
		Overworld: region.NewDimension(path + "/region"),
		path:      path,
	}
	w.Level, err = loadWorldLevel(path + "/level.dat")

	return w, err
}

func (w *World) Dimension(name string) *region.Dimension {
	switch name {
	case "minecraft:overworld", "overworld":
		return w.Overworld
	default:
		return nil
	}
}
