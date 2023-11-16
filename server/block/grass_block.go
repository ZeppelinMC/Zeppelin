package block

import (
	"github.com/dynamitemc/dynamite/server/block/pos"
	"github.com/dynamitemc/dynamite/server/world"
	"github.com/dynamitemc/dynamite/server/world/chunk"
)

type GrassBlock struct {
	Snowy bool
}

func (GrassBlock) EncodedName() string {
	return "minecraft:grass_block"
}

func (g GrassBlock) Properties() map[string]string {
	return map[string]string{
		"snowy": boolstr(g.Snowy),
	}
}

func (GrassBlock) New(_ string, p map[string]string) chunk.Block {
	var a GrassBlock
	if p["snowy"] == "true" {
		a.Snowy = true
	}
	return a
}

func (g GrassBlock) Tick1(p pos.BlockPosition, d *world.Dimension, _ uint) chunk.Block {
	if g.Snowy {
		b := d.Block(p.X(), p.Y()+1, p.Z())
		if _, ok := b.(Snow); !ok {
			g.Snowy = false
		}
	}
	return g
}
