package world

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"math/rand"
	"os"
	"sync"

	"github.com/aimjel/minecraft/nbt"
	"github.com/dynamitemc/dynamite/server/world/anvil"
)

type World struct {
	mu       sync.RWMutex
	nbt      worldData
	Gamemode byte

	overworld *Dimension
	nether    *Dimension
	theEnd    *Dimension
}

func OpenWorld(name string, flat bool) (*World, error) {
	f, err := os.Open(name + "/level.dat")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var wrld World
	if err = loadWorldData(f, &wrld.nbt); err != nil {
		return nil, fmt.Errorf("%v loading world level data", err)
	}

	wrld.overworld = wrld.NewDimension("minecraft:overworld", anvil.NewReader(name+"/region/", name+"/entities/"))
	wrld.nether = wrld.NewDimension("minecraft:the_nether", anvil.NewReader(name+"/DIM-1/region/", name+"/DIM-1/entities/"))
	wrld.theEnd = wrld.NewDimension("minecraft:the_end", anvil.NewReader(name+"/DIM1/region/", name+"/DIM1/entities/"))
	/*if flat {
		wrld.overworld.generator = new(FlatGenerator)
	}*/

	return &wrld, nil
}

func (w *World) Seed() int64 {
	return w.nbt.Data.WorldGenSettings.Seed
}

func (w *World) Spawn() (x, y, z int32, angle float32) {
	return w.nbt.Data.SpawnX, w.nbt.Data.SpawnY, w.nbt.Data.SpawnZ, w.nbt.Data.SpawnAngle
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

func (w *World) GetDimension(typ string) *Dimension {
	switch typ {
	case "minecraft:the_nether":
		return w.Nether()
	case "minecraft:the_end":
		return w.TheEnd()
	default:
		return w.Overworld()
	}
}

func (w *World) LoadSpawnChunks(rd int32) (success int) {
	ow := w.Overworld()
	s := 0
	for x := -rd; x < rd; x++ {
		for z := -rd; z < rd; z++ {
			if _, err := ow.Chunk(x, z); err == nil {
				s++
			}
		}
	}
	return s
}

func (w *World) Gamerules() map[string]string {
	return w.nbt.Data.GameRules
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

func RandomSeed() int64 {
	return rand.Int63()
}

func (w *World) IncrementTime() (worldAge int64, dayTime int64) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.nbt.Data.Time++
	w.nbt.Data.DayTime++
	return w.nbt.Data.Time, w.nbt.Data.DayTime
}
