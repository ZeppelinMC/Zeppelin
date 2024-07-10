package region

import (
	"aether/net/io"
	"aether/net/packet/play"
	"aether/net/registry"
	"aether/server/world/region/blocks"
	"fmt"
)

var emptyLightBuffer = make([]byte, 2048)

func (chunk Chunk) Encode() *play.ChunkDataUpdateLight {
	pk := &play.ChunkDataUpdateLight{
		CX: chunk.XPos,
		CZ: chunk.ZPos,

		Heightmaps: chunk.Heightmaps,

		SkyLightMask:      make(io.BitSet, 1),
		EmptySkyLightMask: make(io.BitSet, 1),
		SkyLightArrays:    make([][]byte, 0, len(chunk.Sections)+2),

		BlockLightMask:      make(io.BitSet, 1),
		EmptyBlockLightMask: make(io.BitSet, 1),
		BlockLightArrays:    make([][]byte, 0, len(chunk.Sections)+2),
	}
	pk.SkyLightArrays = append(pk.SkyLightArrays, emptyLightBuffer)
	pk.SkyLightMask.Set(0, true)
	pk.EmptySkyLightMask.Set(0, true)

	pk.BlockLightArrays = append(pk.BlockLightArrays, emptyLightBuffer)
	pk.BlockLightMask.Set(0, true)
	pk.EmptyBlockLightMask.Set(0, true)

	var data []byte

	for secI, section := range chunk.Sections {
		var blockCount int16
		var airId = -1

		for i, state := range section.BlockStates.Palette {
			if state.Name == "minecraft:air" {
				airId = i
				break
			}
		}
		if airId == -1 {
			blockCount = 4096
		}

		blockBitsPerEntry := byte(len(section.BlockStates.Data) * 64 / 4096)

		if blockCount != 4096 {
			for _, long := range section.BlockStates.Data {
				var pos byte

				for i := 0; i < 64; i++ {
					if blockCount == 4096 {
						break
					}
					if pos+blockBitsPerEntry > 64-pos {
						break
					}

					var entry = (long >> pos) & (int64((1 << blockBitsPerEntry) - 1))

					if entry != int64(airId) {
						blockCount++
					}

					pos += blockBitsPerEntry
				}
			}
		}

		//Block Count
		data = io.AppendShort(data, blockCount)

		//
		// Block Palette
		//

		data = io.AppendUbyte(data, blockBitsPerEntry)

		switch {
		case blockBitsPerEntry == 0:
			pale := section.BlockStates.Palette[0]
			stateId, _ := blocks.Blocks[pale.Name].FindState(pale.Properties)
			data = io.AppendVarInt(data, stateId)
		case blockBitsPerEntry >= 4 && blockBitsPerEntry <= 8:
			data = io.AppendVarInt(data, int32(len(section.BlockStates.Palette)))
			for _, e := range section.BlockStates.Palette {
				stateId, _ := blocks.Blocks[e.Name].FindState(e.Properties)
				data = io.AppendVarInt(data, stateId)
			}
		case blockBitsPerEntry == 15: // no palette
		default:
			fmt.Println("invalid block bits per entry", blockBitsPerEntry, (len(section.BlockStates.Data)*64)/4096)
		}

		data = io.AppendVarInt(data, int32(len(section.BlockStates.Data)))
		for _, long := range section.BlockStates.Data {
			data = io.AppendLong(data, long)
		}

		//
		// Biome Palette
		//

		biomeBitsPerEntry := byte((len(section.Biomes.Data) * 64) / 64)
		data = io.AppendUbyte(data, biomeBitsPerEntry)

		var biomeMap = registry.BiomeId.GetMap()

		switch {
		case biomeBitsPerEntry == 0:
			pale := section.Biomes.Palette[0]
			stateId := biomeMap[pale]

			data = io.AppendVarInt(data, stateId)
		case biomeBitsPerEntry >= 1 && biomeBitsPerEntry <= 3:
			data = io.AppendVarInt(data, int32(len(section.Biomes.Palette)))
			for _, e := range section.Biomes.Palette {
				stateId := biomeMap[e]

				data = io.AppendVarInt(data, stateId)
			}
		case biomeBitsPerEntry == 6: // no palette
		default:
			fmt.Println("invalid biome bits per entry", biomeBitsPerEntry)
		}

		data = io.AppendVarInt(data, int32(len(section.Biomes.Data)))
		for _, long := range section.Biomes.Data {
			data = io.AppendLong(data, long)
		}

		//
		// Lighting
		//

		if section.SkyLight != nil {
			pk.SkyLightMask.Set(secI+1, true)
			if allZero(section.SkyLight) {
				pk.EmptySkyLightMask.Set(secI+1, true)
			}
			pk.SkyLightArrays = append(pk.SkyLightArrays, section.SkyLight)
		}

		if section.BlockLight != nil {
			pk.BlockLightMask.Set(secI+1, true)
			if allZero(section.BlockLight) {
				pk.EmptyBlockLightMask.Set(secI+1, true)
			}
			pk.BlockLightArrays = append(pk.BlockLightArrays, section.BlockLight)
		}
	}
	pk.SkyLightArrays = append(pk.SkyLightArrays, emptyLightBuffer)
	pk.SkyLightMask.Set(len(chunk.Sections), true)
	pk.EmptySkyLightMask.Set(len(chunk.Sections), true)

	pk.BlockLightArrays = append(pk.BlockLightArrays, emptyLightBuffer)
	pk.BlockLightMask.Set(len(chunk.Sections), true)
	pk.EmptyBlockLightMask.Set(len(chunk.Sections), true)

	pk.Data = data

	return pk
}

func allZero(inp []byte) bool {
	for _, i := range inp {
		if i != 0 {
			return false
		}
	}
	return true
}
