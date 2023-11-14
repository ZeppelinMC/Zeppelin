package world

import (
	"sync"

	"github.com/dynamitemc/dynamite/server/world/anvil"
	"github.com/dynamitemc/dynamite/server/world/chunk"
)

type Dimension struct {
	world *World

	typ string

	rd *anvil.Reader

	generator Generator
	//TODO chunk generator and chunk writer

	seed int64

	chunks map[chunk.Hash]*chunk.Chunk

	mu sync.RWMutex
}

func (w *World) NewDimension(typ string, rd *anvil.Reader) *Dimension {
	return &Dimension{
		typ:    typ,
		rd:     rd,
		chunks: make(map[chunk.Hash]*chunk.Chunk),
		world:  w,
	}
}

func (d *Dimension) World() *World {
	return d.world
}

func (d *Dimension) Chunks() map[chunk.Hash]*chunk.Chunk {
	return d.chunks
}

func (d *Dimension) Chunk(x, z int32) (*chunk.Chunk, error) {
	hash := chunk.HashXZ(x, z)

	d.mu.RLock()
	c, ok := d.chunks[hash]
	d.mu.RUnlock()
	if ok {
		return c, nil
	}

	ch, err := d.rd.ReadChunk(x, z)
	if err != nil {
		if d.generator == nil {
			return nil, err
		}
		ch, err = d.generator.GenerateChunk(x, z)
		if err != nil {
			return nil, err
		}
	}
	ents, _ := d.rd.ReadChunkEntities(x, z)
	ch.Entities = ents

	d.mu.Lock()
	defer d.mu.Unlock()
	d.chunks[hash] = ch

	return ch, nil
}

func (d *Dimension) Type() string {
	return d.typ
}

func (d *Dimension) LoadedChunks() int {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return len(d.chunks)
}

func (d *Dimension) Seed() int64 {
	return d.seed
}

func (d *Dimension) Block(x, y, z int64) chunk.Block {
	c, err := d.Chunk(int32(x/16), int32(z/16))
	if err != nil {
		return nil
	}
	b := c.Block(x, y, z)
	if b == nil {
		return chunk.GetBlock("minecraft:air")
	}
	return b
}

func (d *Dimension) SetBlock(x, y, z int64, b chunk.Block) {
	chunk, err := d.Chunk(int32(x/16), int32(z/16))
	if err != nil {
		return
	}
	chunk.SetBlock(x, y, z, b)
}

func ParsePosition(pos int64) (x, y, z int64) {
	x = pos >> 38
	y = pos << 52 >> 52
	z = pos << 26 >> 38
	return
}

type SetBlockHandling = int8

const (
	SetBlockReplace SetBlockHandling = iota
	SetBlockKeep
	SetBlockDestroy
)
