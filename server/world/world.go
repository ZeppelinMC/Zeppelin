package world

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"os"

	"github.com/aimjel/minecraft/nbt"
	"github.com/dynamitemc/dynamite/server/world/anvil"
)

type World struct {
	nbt worldData

	overworld *Dimension
	nether    *Dimension
	theEnd    *Dimension
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
	wrld.overworld = NewDimension("minecraft:overworld", rd)

	return &wrld, nil
}

func (w *World) Seed() int64 {
	return w.nbt.Data.WorldGenSettings.Seed
}

func (w *World) Spawn() (x, y, z int32) {
	return w.nbt.Data.SpawnX, w.nbt.Data.SpawnY, w.nbt.Data.SpawnZ
}

func (w *World) Overworld() *Dimension {
	return w.overworld
}

func (w *World) Nether() *Dimension {
	return w.nether
}

func (w *World) TheEnd() *Dimension {
	return w.theEnd
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
