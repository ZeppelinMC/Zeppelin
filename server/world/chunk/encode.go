package chunk

import (
	"bytes"
	"slices"
	"unsafe"

	"github.com/zeppelinmc/zeppelin/protocol/net/io/buffers"
	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet/play"
	"github.com/zeppelinmc/zeppelin/server/world/chunk/section"
	"github.com/zeppelinmc/zeppelin/util/log"
)

var emptyLightBuffer = make([]byte, 2048)
var fullLightBuffer = make([]byte, 2048)

func init() {
	for i := range fullLightBuffer {
		fullLightBuffer[i] = 0xFF
	}
}

func (chunk *Chunk) Encode(biomeIndexes []string) *play.ChunkDataUpdateLight {
	buf := buffers.Buffers.Get().(*bytes.Buffer)
	buf.Reset()
	defer buffers.Buffers.Put(buf)
	return chunk.EncodeBuf(biomeIndexes, buf)
}

func (chunk *Chunk) EncodeBuf(biomeIndexes []string, buf *bytes.Buffer) *play.ChunkDataUpdateLight {
	bstateId := int32(slices.Index(biomeIndexes, "minecraft:plains"))

	w := encoding.NewWriter(buf)

	pk := &play.ChunkDataUpdateLight{
		CX: chunk.X,
		CZ: chunk.Z,

		Data: buf,

		Heightmaps: *(*play.Heightmaps)(unsafe.Pointer(&chunk.Heightmaps)),

		//BlockEntities: make([]play.BlockEntity, len(chunk.BlockEntities)),

		SkyLightMask:      make(encoding.BitSet, 1),
		EmptySkyLightMask: make(encoding.BitSet, 1),
		SkyLightArrays:    make([][]byte, 1, len(chunk.Sections)+1),

		BlockLightMask:      make(encoding.BitSet, 1),
		EmptyBlockLightMask: make(encoding.BitSet, 1),
		BlockLightArrays:    make([][]byte, 1, len(chunk.Sections)+1),
	}

	pk.SkyLightMask.Set(1)
	pk.EmptySkyLightMask.Set(1)
	pk.SkyLightArrays[0] = emptyLightBuffer

	pk.BlockLightMask.Set(1)
	pk.EmptyBlockLightMask.Set(1)
	pk.BlockLightArrays[0] = emptyLightBuffer

	/*for i, entity := range chunk.BlockEntities {
		pk.BlockEntities[i] = play.BlockEntity{
			X:    entity.X,
			Y:    entity.Y,
			Z:    entity.Z,
			Type: registry.BlockEntityType.Get(entity.Id),
			Data: entity,
		}
	}*/

	for secI, sec := range chunk.Sections {
		var blockCount int16 = 1024
		blockBitsPerEntry, biomeBitsPerEntry := sec.BPE()
		//var airId = -1

		/*for i, state := range blockPalette {
			name, _ := state.Encode()
			if name == "minecraft:air" {
				airId = i
				break
			}
		}
		if airId == -1 {
			blockCount = 4096
		}

		if blockCount != 4096 {
			for _, long := range blockStates {
				var pos int

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
		}*/

		//Block Count
		w.Short(blockCount)

		//
		// Block Palette
		//
		w.Ubyte(byte(blockBitsPerEntry))

		switch {
		case blockBitsPerEntry == 0:
			stateId, _ := section.BlockStateId(sec.BlockStates.Palette[0])
			w.VarInt(stateId)
		case blockBitsPerEntry >= 4 && blockBitsPerEntry <= 8:
			w.VarInt(int32(len(sec.BlockStates.Palette)))
			for _, e := range sec.BlockStates.Palette {
				stateId, _ := section.BlockStateId(e)
				w.VarInt(stateId)
			}
		case blockBitsPerEntry == 15: // no palette
		default:
			log.Println("invalid block bits per entry", blockBitsPerEntry, (len(sec.BlockStates.Data)*64)/4096)
		}

		w.VarInt(int32(len(sec.BlockStates.Data)))
		for _, long := range sec.BlockStates.Data {
			w.Long(long)
		}

		//
		// Biome Palette
		//

		w.Ubyte(byte(biomeBitsPerEntry))

		switch {
		case biomeBitsPerEntry == 0:
			//pale := biomePalette[0]
			/*stateId := int32(slices.Index(biomeIndexes, pale))
			if stateId == -1 {
				fmt.Println("h", pale)
			}*/

			w.VarInt(bstateId)
		case biomeBitsPerEntry >= 1 && biomeBitsPerEntry <= 3:
			w.VarInt(int32(len(sec.Biomes.Palette)))
			for _, e := range sec.Biomes.Palette {
				_ = e
				/*stateId := int32(slices.Index(biomeIndexes, e))
				if stateId == -1 {
					fmt.Println("h", e)
				}*/

				w.VarInt(bstateId)
			}
		case biomeBitsPerEntry == 6: // no palette
		default:
			log.Println("invalid biome bits per entry", pk.CX, pk.CZ, sec.Y, biomeBitsPerEntry)
		}

		w.VarInt(int32(len(sec.Biomes.Data)))
		for _, long := range sec.Biomes.Data {
			w.Long(long)
		}

		//
		// Lighting
		//

		if sec.SkyLight != nil {
			pk.SkyLightMask.Set(secI + 1)
			if allZero(sec.SkyLight) {
				pk.EmptySkyLightMask.Set(secI + 1)
			}
			pk.SkyLightArrays = append(pk.SkyLightArrays, *(*[]byte)(unsafe.Pointer(&sec.SkyLight)))
		}

		if sec.BlockLight != nil {
			pk.BlockLightMask.Set(secI + 1)
			if allZero(sec.BlockLight) {
				pk.EmptyBlockLightMask.Set(secI + 1)
			}
			pk.BlockLightArrays = append(pk.BlockLightArrays, *(*[]byte)(unsafe.Pointer(&sec.BlockLight)))
		}
	}
	/*pk.SkyLightArrays = append(pk.SkyLightArrays, emptyLightBuffer)
	pk.SkyLightMask.Set(len(chunk.sections))
	pk.EmptySkyLightMask.Set(len(chunk.sections))

	pk.BlockLightArrays = append(pk.BlockLightArrays, emptyLightBuffer)
	pk.BlockLightMask.Set(len(chunk.sections))
	pk.EmptyBlockLightMask.Set(len(chunk.sections))*/

	//for i := 0; i < 24; i++ {
	//	fmt.Println(pk.SkyLightMask.Get(i), len(pk.SkyLightArrays[i+1]))
	//}

	//pk.Data = buf.Bytes()

	return pk
}

func allZero(inp []int8) bool {
	for _, i := range inp {
		if i != 0 {
			return false
		}
	}
	return true
}
