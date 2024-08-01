package block

import (
	"strconv"

	"github.com/zeppelinmc/zeppelin/net/packet/play"
	"github.com/zeppelinmc/zeppelin/server/session"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
	"github.com/zeppelinmc/zeppelin/server/world/chunk/section"
	"github.com/zeppelinmc/zeppelin/server/world/dimension"
)

type BlockProperties = map[string]string
type Block = section.Block

type Axis = string

const (
	AxisX Axis = "x"
	AxisY Axis = "y"
	AxisZ Axis = "z"
)

type Direction = string

const (
	DirectionNorth Direction = "north"
	DirectionSouth Direction = "south"
	DirectionWest  Direction = "west"
	DirectionEast  Direction = "east"
)

func atoi(str string) int {
	v, _ := strconv.Atoi(str)

	return v
}

// A usable block is a block which performs a certain action when left clicked
type Usable interface {
	Block
	Use(clicker session.Session, pk play.UseItemOn, dimension *dimension.Dimension)
}

// A block entity haver is one that also has a block entity
type BlockEntityHaver interface {
	Block
	BlockEntity() chunk.BlockEntity
}

func init() {
	section.RegisterBlock(Air{})
	section.RegisterBlock(Bedrock{})
	section.RegisterBlock(Dirt{})
	section.RegisterBlock(GrassBlock{})
	section.RegisterBlock(OakLog{})
	section.RegisterBlock(Chest{})
}
