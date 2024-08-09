package world

import (
	"strings"

	"github.com/zeppelinmc/zeppelin/atomic"
	"github.com/zeppelinmc/zeppelin/log"
	"github.com/zeppelinmc/zeppelin/server/session"
	"github.com/zeppelinmc/zeppelin/server/world/dimension"
	"github.com/zeppelinmc/zeppelin/server/world/level"
	"github.com/zeppelinmc/zeppelin/server/world/terrain"
)

type World struct {
	level.Level
	dimensions map[string]*dimension.Dimension
	Broadcast  *session.Broadcast

	path              string
	worldAge, dayTime atomic.AtomicValue[int64]
}

func NewWorld(path string) (*World, error) {
	var err error
	w := &World{
		path:      path,
		Broadcast: session.NewBroadcast(),
	}
	w.Level, err = level.LoadWorldLevel(path)
	w.worldAge = atomic.Value(w.Level.Data.Time)
	w.dayTime = atomic.Value(w.Level.Data.DayTime)
	w.dimensions = map[string]*dimension.Dimension{
		"minecraft:overworld": dimension.New(
			path+"/region",
			"minecraft:overworld",
			"minecraft:overworld",
			w.Broadcast,
			terrain.NewTerrainGenerator(int64(w.Data.WorldGenSettings.Seed)),
			true,
		),
	}

	return w, err
}

// returns the dimension struct for the dimension name
func (w *World) Dimension(name string) *dimension.Dimension {
	if !strings.Contains(name, ":") {
		name = "minecraft:" + name
	}

	return w.dimensions[name]
}

func (w *World) Save() {
	for _, dim := range w.dimensions {
		dim.Save()
		log.Infoln("Saved dimension", dim.Name())
	}
}

func (w *World) RegisterDimension(name string, dim *dimension.Dimension) {
	w.dimensions[name] = dim
}

// increments the day time and world age by one tick and returns the updated time
func (w *World) IncrementTime() (worldAge, dayTime int64) {
	worldAge = w.worldAge.Get() + 1
	dayTime = w.dayTime.Get() + 1

	w.worldAge.Set(worldAge)
	w.dayTime.Set(dayTime)

	return
}

func (w *World) LoadedChunks() int32 {
	var count int32

	for _, dim := range w.dimensions {
		count += dim.LoadedChunks()
	}

	return count
}
