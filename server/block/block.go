package block

import (
	"github.com/dynamitemc/dynamite/server/block/pos"
	"github.com/dynamitemc/dynamite/server/world"
	"github.com/dynamitemc/dynamite/server/world/chunk"
)

func init() {
	chunk.RegisterBlock(Air{})
	chunk.RegisterBlock(Dirt{})
	chunk.RegisterBlock(Dirt{true})
	chunk.RegisterBlock(GrassBlock{})
	chunk.RegisterBlock(Snow{})
	chunk.RegisterBlock(Bedrock{})
	chunk.RegisterBlock(AcaciaLog{})
	chunk.RegisterBlock(AcaciaLog{Stripped: true})

	chunk.RegisterBlock(BirchLog{})
	chunk.RegisterBlock(BirchLog{Stripped: true})

	chunk.RegisterBlock(CherryLog{})
	chunk.RegisterBlock(CherryLog{Stripped: true})

	chunk.RegisterBlock(DarkOakLog{})
	chunk.RegisterBlock(DarkOakLog{Stripped: true})

	chunk.RegisterBlock(OakLog{})
	chunk.RegisterBlock(OakLog{Stripped: true})

	chunk.RegisterBlock(SpruceLog{})
	chunk.RegisterBlock(SpruceLog{Stripped: true})
}

type BreakInfo struct {
	Hardness    float64
	Unbreakable bool
}

type Ticker interface {
	chunk.Block
	Tick(pos.BlockPosition, *world.Dimension, uint) chunk.Block
}

type RandomTicker interface {
	chunk.Block
	RandomTick(pos.BlockPosition, *world.Dimension, uint) chunk.Block
}

type Breakable interface {
	chunk.Block
	BreakInfo() BreakInfo
}

func boolstr(b bool) (s string) {
	s = "false"
	if b {
		s = "true"
	}
	return
}
