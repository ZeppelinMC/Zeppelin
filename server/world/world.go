package world

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/aimjel/minecraft/nbt"
	"github.com/dynamitemc/dynamite/server/world/anvil"
	"os"
)

type worldData struct {
	Data struct {
		WorldGenSettings struct {
			Seed int64 `nbt:"seed"`
		}
		DataVersion int32
		GameRules   map[string]interface{}
	}
}

type World struct {
	nbt worldData

	dimensions []*Dimension
}

func OpenWorld(name string) (*World, error) {
	f, err := os.Open(name + "/level.dat")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var wrld World
	if err = loadWorldData(f, &wrld.nbt); err != nil {
		return nil, fmt.Errorf("%v loading world level data", err)
	}

	rd := anvil.NewReader(name + "/region/")
	ow := NewDimension("minecraft:overworld", rd)
	wrld.dimensions = append(wrld.dimensions, ow)

	return &wrld, nil
}

func (w *World) Seed() int64 {
	return w.nbt.Data.WorldGenSettings.Seed
}

func (w *World) DefaultDimension() *Dimension {
	return w.dimensions[0]
}

func loadWorldData(f *os.File, wNbt *worldData) error {
	gzipRd, err := gzip.NewReader(f)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if _, err := buf.ReadFrom(gzipRd); err != nil {
		return err
	}

	return nbt.Unmarshal(buf.Bytes(), wNbt)
}
