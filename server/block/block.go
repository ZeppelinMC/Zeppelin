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
}

type BreakInfo struct {
	Unbreakable bool
}

type Ticker interface {
	chunk.Block
	Tick(pos.BlockPosition, *world.Dimension, uint) chunk.Chunk
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
