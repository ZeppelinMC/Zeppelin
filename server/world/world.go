package world

import (
	"math"
	"strings"

	"github.com/zeppelinmc/zeppelin/atomic"
	"github.com/zeppelinmc/zeppelin/server/world/region"
	"github.com/zeppelinmc/zeppelin/server/world/terrain"
)

type World struct {
	Level
	dimensions map[string]*region.Dimension

	path              string
	worldAge, dayTime atomic.AtomicValue[int64]
}

func NewWorld(path string) (*World, error) {
	var err error
	w := &World{
		path: path,
	}
	w.Level, err = loadWorldLevel(path + "/level.dat")
	w.worldAge = atomic.Value(w.Level.Data.Time)
	w.dayTime = atomic.Value(w.Level.Data.DayTime)
	w.dimensions = map[string]*region.Dimension{
		"minecraft:overworld": region.NewDimension(path+"/region", 0, terrain.NewTerrainGenerator(int64(w.Data.WorldGenSettings.Seed))),
	}

	return w, err
}

// returns the dimension struct for the dimension name
func (w *World) Dimension(name string) *region.Dimension {
	if !strings.Contains(name, ":") {
		name = "minecraft:" + name
	}

	return w.dimensions[name]
}

func (w *World) RegisterDimension(name string, dim *region.Dimension) {
	w.dimensions[name] = dim
}

// increments the day time and world age by one tick and returns the updated time
func (w *World) IncrementTime() (worldAge, dayTime int64) {
	worldAge = w.worldAge.Get() + 1
	dayTime = (w.dayTime.Get() + 1) % 24000

	w.worldAge.Set(worldAge)
	w.dayTime.Set(dayTime)

	return
}

// HashCode is an implementation of Java's hashCode function. It used to turn any string seed into a long seed
func HashCode(s string) int64 {
	var result int64
	n := len(s)

	for i := 0; i < len(s)-1; i++ {
		result += int64(s[i]) * int64(math.Pow(31, float64(n-(i+1))))
	}

	return result + int64(s[int(n)-1])
}
