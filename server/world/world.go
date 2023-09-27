package world

import (
	"bytes"
	"compress/gzip"
	"os"
	"sync/atomic"

	"github.com/aimjel/minecraft/nbt"
)

type worldData struct {
	Data struct {
		WorldGenSettings struct {
			Seed int64 `nbt:"seed"`
		}
		DataVersion int32
		GameRules   map[string]string
	}
}

type World struct {
	nbt worldData

	entityIdCounter atomic.Value

	dimensions []*Dimension
}

func OpenWorld(name string) (*World, error) {
	f, err := os.Open(name + "/level.dat")
	if err != nil {
		return nil, err
	}

	gzipRd, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if _, err := buf.ReadFrom(gzipRd); err != nil {
		return nil, err
	}

	var wrld World
	if err := nbt.Unmarshal(buf.Bytes(), &wrld.nbt); err != nil {
		return nil, err
	}

	//todo temp
	wrld.dimensions = make([]*Dimension, 0, 1)
	wrld.dimensions = append(wrld.dimensions, NewDimension("minecraft:overworld"))

	return &wrld, nil
}

func (w *World) Seed() int64 {
	return w.nbt.Data.WorldGenSettings.Seed
}

func (w *World) DefaultDimension() *Dimension {
	return w.dimensions[0]
}
