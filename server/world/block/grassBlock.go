package block

import (
	"strconv"
)

type GrassBlock struct {
	Snowy bool
}

func (g GrassBlock) Encode() (string, BlockProperties) {
	return "minecraft:grass_block", BlockProperties{
		"snowy": strconv.FormatBool(g.Snowy),
	}
}

func (g GrassBlock) New(props BlockProperties) Block {
	return GrassBlock{Snowy: props["snowy"] == "true"}
}

/*func (g GrassBlock) PlaceSound(pos pos.BlockPosition) *play.SoundEffect {
	return session.SoundEffect("minecraft:block.grass.place", false, nil, play.SoundCategoryBlock, pos.X(), pos.Y(), pos.Z(), 1, 1)
}
*/
