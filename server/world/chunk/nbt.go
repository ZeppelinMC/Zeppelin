package chunk

type anvilChunk struct {
	Status string

	XPos int32 `nbt:"xPos"`
	YPos int32 `nbt:"yPos"`
	ZPos int32 `nbt:"zPos"`

	Heightmaps HeightMap

	Sections []anvilChunkSection `nbt:"sections"`

	DataVersion int32
}

type anvilChunkSection struct {
	Y           int8
	BlockStates struct {
		Data    []int64      `nbt:"data"`
		Palette []blockEntry `nbt:"palette"`
	} `nbt:"block_states"`
	Biomes struct {
		Palette []string `nbt:"palette"`
	} `nbt:"biomes"`
	BlockLight []int8
	SkyLight   []int8
}

type blockEntry struct {
	Properties map[string]string
	Name       string
}
