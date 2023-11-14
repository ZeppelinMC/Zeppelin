package block

import (
	"github.com/dynamitemc/dynamite/server/world/chunk"
)

type Dirt struct {
	Coarse bool
}

func (d Dirt) EncodedName() string {
	if d.Coarse {
		return "minecraft:coarse"
	}
	return "minecraft:dirt"
}

func (g Dirt) Properties() map[string]string {
	return nil
}

func (d Dirt) New(string, map[string]string) chunk.Block {
	return d
}
