package dimension

import (
	"github.com/zeppelinmc/zeppelin/util/log"
	"strings"
)

func NewDimensionManager(dimensions map[string]*Dimension) DimensionManager {
	return DimensionManager{dimensions: dimensions}
}

type DimensionManager struct {
	dimensions map[string]*Dimension
}

// Dimension returns the dimension struct for the dimension name
func (w *DimensionManager) Dimension(name string) *Dimension {
	if !strings.Contains(name, ":") {
		name = "minecraft:" + name
	}

	return w.dimensions[name]
}

func (w *DimensionManager) Save() {
	for _, dim := range w.dimensions {
		dim.Save()
	}
	log.Infoln("Closed world")
}

func (w *DimensionManager) RegisterDimension(name string, dim *Dimension) {
	w.dimensions[name] = dim
}

func (w *DimensionManager) LoadedChunks() int32 {
	var count int32

	for _, dim := range w.dimensions {
		count += dim.LoadedChunks()
	}

	return count
}
