package block

import "github.com/dynamitemc/dynamite/server/world/chunk"

type Air struct{}

func (Air) EncodedName() string {
	return "minecraft:air"
}

func (Air) Properties() map[string]string {
	return nil
}

func (a Air) New(string, map[string]string) chunk.Block {
	return a
}
