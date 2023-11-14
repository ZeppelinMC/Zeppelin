package block

import (
	"strconv"

	"github.com/dynamitemc/dynamite/server/world/chunk"
)

type Snow struct {
	Layers int
}

func (Snow) EncodedName() string {
	return "minecraft:snow"
}

func (s Snow) Properties() map[string]string {
	return map[string]string{
		"layers": strconv.Itoa(s.Layers),
	}
}

func (Snow) New(_ string, m map[string]string) chunk.Block {
	l, _ := strconv.Atoi(m["layers"])
	if l == 0 {
		l = 1
	}
	return Snow{
		Layers: l,
	}
}
