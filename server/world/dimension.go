package world

import (
	"errors"
	"github.com/aimjel/nitrate/server/world/anvil"
	"github.com/aimjel/nitrate/server/world/chunk"
	"sync"
)

type Dimension struct {
	// _type of dimension, possible values are:
	// minecraft:overworld
	// minecraft:the_end
	// minecraft:the_nether
	_type string

	chunks map[uint64]*chunk.Chunk

	rd *anvil.Reader

	gen Generator

	mu sync.RWMutex
}

func NewDimension(typ string) *Dimension {
	return &Dimension{
		chunks: make(map[uint64]*chunk.Chunk),
		_type:  typ,
	}
}

func (d *Dimension) Type() string {
	return d._type
}

func (d *Dimension) Chunk(x, z int32) (*chunk.Chunk, error) {
	h := chunk.Hash(x, z)

	d.mu.RLock()
	c, ok := d.chunks[h]
	d.mu.RUnlock()

	if !ok {
		ch, err := d.rd.ReadChunk(x, z)
		if err != nil {
			if !errors.Is(err, chunk.ErrNotFound) && d.gen == nil {
				return nil, err
			}

			ch = d.gen.GenerateChunk(x, z)
		}

		d.mu.Lock()
		d.chunks[h] = ch
		d.mu.Unlock()

		return ch, nil
	}

	return c, nil
}
