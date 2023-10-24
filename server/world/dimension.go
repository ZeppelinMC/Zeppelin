package world

import (
	"sync"

	"github.com/dynamitemc/dynamite/server/world/anvil"
	"github.com/dynamitemc/dynamite/server/world/chunk"
)

type Dimension struct {
	typ string

	rd *anvil.Reader

	generator Generator
	//TODO chunk generator and chunk writer

	seed int64

	chunks map[uint64]*chunk.Chunk

	mu sync.RWMutex
}

func NewDimension(typ string, rd *anvil.Reader) *Dimension {
	return &Dimension{
		typ:    typ,
		rd:     rd,
		chunks: make(map[uint64]*chunk.Chunk),
	}
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
		ch, err = d.generator.Generate(x, z)
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

func (d *Dimension) Block(x, y, z int64) {
	chunk, err := d.Chunk(int32(x/16), int32(z/16))
	if err != nil {
		return
	}
	chunk.Block(x, y, z)
}

func ParsePosition(p uint64) (x, y, z int64) {
	pos := int64(p)
	x = pos >> 38
	y = pos << 52 >> 52
	z = pos << 26 >> 38
	return
}
