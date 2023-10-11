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

	d.mu.Lock()
	defer d.mu.Unlock()
	d.chunks[hash] = ch

	return ch, nil
}

func (d *Dimension) Type() string {
	return d.typ
}

func (d *Dimension) Seed() int64 {
	return d.seed
}
