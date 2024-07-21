package region

import (
	"bytes"
	"fmt"
	"math/bits"

	"github.com/zeppelinmc/zeppelin/net/io"
	"github.com/zeppelinmc/zeppelin/net/packet/play"
	"github.com/zeppelinmc/zeppelin/net/registry"
	"github.com/zeppelinmc/zeppelin/server/world/region/blocks"
)

var emptyLightBuffer = make([]byte, 2048)
var fullLightBuffer = make([]byte, 2048)

func init() {
	for i := range fullLightBuffer {
		fullLightBuffer[i] = 0xFF
	}
}

func (chunk Chunk) Encode(buffer *bytes.Buffer) *play.ChunkDataUpdateLight {
	w := io.NewWriter(buffer)
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
	pk.SkyLightMask.Set(0)
	pk.EmptySkyLightMask.Set(0)

	pk.BlockLightArrays = append(pk.BlockLightArrays, emptyLightBuffer)
	pk.BlockLightMask.Set(0)
	pk.EmptyBlockLightMask.Set(0)

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
		w.Short(blockCount)

		//
		// Block Palette
		//
		w.Ubyte(blockBitsPerEntry)

		switch {
		case blockBitsPerEntry == 0:
			pale := section.BlockStates.Palette[0]
			stateId, _ := blocks.Blocks[pale.Name].FindState(pale.Properties)
			w.VarInt(stateId)
		case blockBitsPerEntry >= 4 && blockBitsPerEntry <= 8:
			w.VarInt(int32(len(section.BlockStates.Palette)))
			for _, e := range section.BlockStates.Palette {
				stateId, _ := blocks.Blocks[e.Name].FindState(e.Properties)
				w.VarInt(stateId)
			}
		case blockBitsPerEntry == 15: // no palette
		default:
			fmt.Println("invalid block bits per entry", blockBitsPerEntry, (len(section.BlockStates.Data)*64)/4096)
		}

		w.VarInt(int32(len(section.BlockStates.Data)))
		for _, long := range section.BlockStates.Data {
			w.Long(long)
		}

		//
		// Biome Palette
		//

		biomeBitsPerEntry := byte(bits.Len32(uint32(len(section.Biomes.Palette))))
		w.Ubyte(biomeBitsPerEntry)

		var biomeMap = registry.BiomeId.GetMap()

		switch {
		case biomeBitsPerEntry == 0:
			pale := section.Biomes.Palette[0]
			stateId := biomeMap[pale]

			w.VarInt(stateId)
		case biomeBitsPerEntry >= 1 && biomeBitsPerEntry <= 3:
			w.VarInt(int32(len(section.Biomes.Palette)))
			for _, e := range section.Biomes.Palette {
				stateId := biomeMap[e]

				w.VarInt(stateId)
			}
		case biomeBitsPerEntry == 6: // no palette
		default:
			fmt.Println("invalid biome bits per entry", pk.CX, pk.CZ, section.Y, biomeBitsPerEntry)
		}

		w.VarInt(int32(len(section.Biomes.Data)))
		for _, long := range section.Biomes.Data {

			w.Long(long)
		}

		//
		// Lighting
		//

		if section.SkyLight != nil {
			pk.SkyLightMask.Set(secI + 1)
			if allZero(section.SkyLight) {
				pk.EmptySkyLightMask.Set(secI + 1)
			}
			pk.SkyLightArrays = append(pk.SkyLightArrays, section.SkyLight)
		}

		if section.BlockLight != nil {
			pk.BlockLightMask.Set(secI + 1)
			if allZero(section.BlockLight) {
				pk.EmptyBlockLightMask.Set(secI + 1)
			}
			pk.BlockLightArrays = append(pk.BlockLightArrays, section.BlockLight)
		}
	}
	pk.SkyLightArrays = append(pk.SkyLightArrays, emptyLightBuffer)
	pk.SkyLightMask.Set(len(chunk.Sections))
	pk.EmptySkyLightMask.Set(len(chunk.Sections))

	pk.BlockLightArrays = append(pk.BlockLightArrays, emptyLightBuffer)
	pk.BlockLightMask.Set(len(chunk.Sections))
	pk.EmptyBlockLightMask.Set(len(chunk.Sections))

	pk.Data = buffer.Bytes()

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
