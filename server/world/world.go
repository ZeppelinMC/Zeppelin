package world

import "aether/server/world/region"

type World struct {
	save *region.Save
}

func NewWorld(path string) *World {
	return &World{
		save: region.NewSave(path + "/region"),
	}
}

func (w *World) GetChunk(x, z int32) (*region.Chunk, error) {
	return w.save.GetChunk(x, z)
}
