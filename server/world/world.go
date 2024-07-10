package world

import "aether/server/world/region"

type World struct {
	save *region.Dimension
}

func NewWorld(path string) *World {
	return &World{
		save: region.NewDimension(path + "/region"),
	}
}

func (w *World) GetChunk(x, z int32) (*region.Chunk, error) {
	return w.save.GetChunk(x, z)
}
