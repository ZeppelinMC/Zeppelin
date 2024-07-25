package region

import (
	"bytes"
	"fmt"
	"math/bits"
	"slices"

	"github.com/zeppelinmc/zeppelin/net/io"
	"github.com/zeppelinmc/zeppelin/net/packet/play"
	"github.com/zeppelinmc/zeppelin/server/world/region/block"
)

var emptyLightBuffer = make([]byte, 2048)
var fullLightBuffer = make([]byte, 2048)

func (chunk Chunk) Encode(biomeIndexes []string) *play.ChunkDataUpdateLight {
	buf := buffers.Get().(*bytes.Buffer)
	buf.Reset()
	defer buffers.Put(buf)

	w := io.NewWriter(buf)
	pk := &play.ChunkDataUpdateLight{
		CX: chunk.X,
		CZ: chunk.Z,

		Heightmaps: chunk.Heightmaps,

		SkyLightMask:      make(io.BitSet, 1),
		EmptySkyLightMask: make(io.BitSet, 1),
		SkyLightArrays:    make([][]byte, 0, len(chunk.sections)+2),

		BlockLightMask:      make(io.BitSet, 1),
		EmptyBlockLightMask: make(io.BitSet, 1),
		BlockLightArrays:    make([][]byte, 0, len(chunk.sections)+2),
	}
	pk.SkyLightArrays = append(pk.SkyLightArrays, emptyLightBuffer)
	pk.SkyLightMask.Set(0)
	pk.EmptySkyLightMask.Set(0)

	pk.BlockLightArrays = append(pk.BlockLightArrays, emptyLightBuffer)
	pk.BlockLightMask.Set(0)
	pk.EmptyBlockLightMask.Set(0)

	for secI, section := range chunk.sections {
		var blockCount int16
		var airId = -1

		for i, state := range section.blockPalette {
			if state.Name == "minecraft:air" {
				airId = i
				break
			}
		}
		if airId == -1 {
			blockCount = 4096
		}

		if blockCount != 4096 {
			for _, long := range section.blockStates {
				var pos int

				for i := 0; i < 64; i++ {
					if blockCount == 4096 {
						break
					}
					if pos+section.blockBitsPerEntry > 64-pos {
						break
					}

					var entry = (long >> pos) & (int64((1 << section.blockBitsPerEntry) - 1))

					if entry != int64(airId) {
						blockCount++
					}

					pos += section.blockBitsPerEntry
				}
			}
		}

		//Block Count
		w.Short(blockCount)

		//
		// Block Palette
		//
		w.Ubyte(byte(section.blockBitsPerEntry))

		switch {
		case section.blockBitsPerEntry == 0:
			pale := section.blockPalette[0]
			stateId, _ := block.Blocks[pale.Name].FindState(pale.Properties)
			w.VarInt(stateId)
		case section.blockBitsPerEntry >= 4 && section.blockBitsPerEntry <= 8:
			w.VarInt(int32(len(section.blockPalette)))
			for _, e := range section.blockPalette {
				stateId, _ := block.Blocks[e.Name].FindState(e.Properties)
				w.VarInt(stateId)
			}
		case section.blockBitsPerEntry == 15: // no palette
		default:
			fmt.Println("invalid block bits per entry", section.blockBitsPerEntry, (len(section.blockStates)*64)/4096)
		}

		w.VarInt(int32(len(section.blockStates)))
		for _, long := range section.blockStates {
			w.Long(long)
		}

		//
		// Biome Palette
		//

		biomeBitsPerEntry := byte(bits.Len32(uint32(len(section.biomes.Palette))))
		w.Ubyte(biomeBitsPerEntry)

		switch {
		case biomeBitsPerEntry == 0:
			pale := section.biomes.Palette[0]
			stateId := int32(slices.Index(biomeIndexes, pale))

			w.VarInt(stateId)
		case biomeBitsPerEntry >= 1 && biomeBitsPerEntry <= 3:
			w.VarInt(int32(len(section.biomes.Palette)))
			for _, e := range section.biomes.Palette {
				stateId := int32(slices.Index(biomeIndexes, e))

				w.VarInt(stateId)
			}
		case biomeBitsPerEntry == 6: // no palette
		default:
			fmt.Println("invalid biome bits per entry", pk.CX, pk.CZ, section.y, biomeBitsPerEntry)
		}

		w.VarInt(int32(len(section.biomes.Data)))
		for _, long := range section.biomes.Data {
			w.Long(long)
		}

		//
		// Lighting
		//

		if section.skyLight != nil {
			pk.SkyLightMask.Set(secI + 1)
			if allZero(section.skyLight) {
				pk.EmptySkyLightMask.Set(secI + 1)
			}
			pk.SkyLightArrays = append(pk.SkyLightArrays, section.skyLight)
		}

		if section.blockLight != nil {
			pk.BlockLightMask.Set(secI + 1)
			if allZero(section.blockLight) {
				pk.EmptyBlockLightMask.Set(secI + 1)
			}
			pk.BlockLightArrays = append(pk.BlockLightArrays, section.blockLight)
		}
	}
	pk.SkyLightArrays = append(pk.SkyLightArrays, emptyLightBuffer)
	pk.SkyLightMask.Set(len(chunk.sections))
	pk.EmptySkyLightMask.Set(len(chunk.sections))

	pk.BlockLightArrays = append(pk.BlockLightArrays, emptyLightBuffer)
	pk.BlockLightMask.Set(len(chunk.sections))
	pk.EmptyBlockLightMask.Set(len(chunk.sections))

	pk.Data = buf.Bytes()

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
