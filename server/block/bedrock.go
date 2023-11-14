package block

import "github.com/dynamitemc/dynamite/server/world/chunk"

type Bedrock struct{}

func (d Bedrock) BreakInfo() BreakInfo {
	return BreakInfo{
		Unbreakable: true,
	}
}

func (d Bedrock) EncodedName() string {
	return "minecraft:bedrock"
}

func (g Bedrock) Properties() map[string]string {
	return nil
}

func (d Bedrock) New(p map[string]string) chunk.Block {
	return d
}
