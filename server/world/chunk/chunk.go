package chunk

import (
	"errors"
)

var ErrNotFound = errors.New("chunk not found")

type Chunk struct {
}

func NewAnvilChunk(b []byte) (*Chunk, error) {
	return nil, nil
}
