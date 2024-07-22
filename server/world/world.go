package world

import (
	"github.com/zeppelinmc/zeppelin/atomic"
	"github.com/zeppelinmc/zeppelin/server/world/region"
)

type World struct {
	Level
	Overworld *region.Dimension

	path              string
	worldAge, dayTime atomic.AtomicValue[int64]
}

func NewWorld(path string) (*World, error) {
	var err error
	w := &World{
		Overworld: region.NewDimension(path+"/region", 0),
		path:      path,
	}
	w.Level, err = loadWorldLevel(path + "/level.dat")
	w.worldAge = atomic.Value(w.Level.Data.Time)
	w.dayTime = atomic.Value(w.Level.Data.DayTime)

	return w, err
}

// returns the dimension struct for the dimension name
func (w *World) Dimension(name string) *region.Dimension {
	switch name {
	case "minecraft:overworld", "overworld":
		return w.Overworld
	default:
		return nil
	}
}

// increments the day time and world age by one tick and returns the updated time
func (w *World) IncrementTime() (worldAge, dayTime int64) {
	worldAge = w.worldAge.Get() + 1
	dayTime = (w.dayTime.Get() + 1) % 24000

	w.worldAge.Set(worldAge)
	w.dayTime.Set(dayTime)

	return
}
