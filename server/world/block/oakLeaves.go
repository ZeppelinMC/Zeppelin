package block

import (
	"fmt"
	"strconv"

	"github.com/zeppelinmc/zeppelin/net/packet/play"
	"github.com/zeppelinmc/zeppelin/server/session"
	"github.com/zeppelinmc/zeppelin/text"
)

type OakLeaves struct {
	Distance    int
	Persistent  bool
	Waterlogged bool
}

func (g OakLeaves) Encode() (string, BlockProperties) {
	return "minecraft:oak_leaves", BlockProperties{
		"distance":    strconv.Itoa(g.Distance),
		"persistent":  strconv.FormatBool(g.Persistent),
		"waterlogged": strconv.FormatBool(g.Waterlogged),
	}
}

func (g OakLeaves) New(props BlockProperties) Block {
	return OakLeaves{
		Distance:    atoi(props["distance"]),
		Persistent:  props["persistent"] == "true",
		Waterlogged: props["waterlogged"] == "true",
	}
}

func (g OakLeaves) Use(clicker session.Session, pk play.UseItemOn) {
	clicker.SystemMessage(text.TextComponent{Text: fmt.Sprintf("kill yourself (at %d %d %d)", pk.BlockX, pk.BlockY, pk.BlockZ)})
}

var _ Block = (*OakLeaves)(nil)
