package player

import (
	"bytes"
	"github.com/zeppelinmc/zeppelin/protocol/net/io/buffers"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet/play"
	"github.com/zeppelinmc/zeppelin/server/world/dimension"
)

// Dimension returns the dimension struct this player is in
func (p *Player) Dimension() *dimension.Dimension {
	return p.dimensionManager.Dimension(p.DimensionName())
}

func (p *Player) sendSpawnChunks() error {
	viewDistance := p.ViewDistance()

	x, _, z := p.Position()
	chunkX, chunkZ := int32(x)>>4, int32(z)>>4

	if err := p.WritePacket(&play.SetCenterChunk{ChunkX: chunkX, ChunkZ: chunkZ}); err != nil {
		return err
	}

	var chunks int32

	if err := p.WritePacket(&play.ChunkBatchStart{}); err != nil {
		return err
	}

	buf := buffers.Buffers.Get().(*bytes.Buffer)

	for x := chunkX - viewDistance; x < chunkX+viewDistance; x++ {
		for z := chunkZ - viewDistance; z < chunkZ+viewDistance; z++ {
			buf.Reset()
			c, err := p.Dimension().GetChunkBuf(x, z, buf)
			if err != nil {
				continue
			}
			buf.Reset()

			if err := p.WritePacket(c.EncodeBuf(p.registryIndexes["minecraft:worldgen/biome"], buf)); err != nil {
				return err
			}
			chunks++
		}
	}
	buffers.Buffers.Put(buf)

	if err := p.WritePacket(&play.ChunkBatchFinished{
		BatchSize: chunks,
	}); err != nil {
		return err
	}

	_, err := p.awaitPacket(play.PacketIdChunkBatchReceived)

	return err
}
