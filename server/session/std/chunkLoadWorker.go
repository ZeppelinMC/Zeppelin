package std

import (
	"github.com/zeppelinmc/zeppelin/protocol/net/packet/play"
	"github.com/zeppelinmc/zeppelin/server/world/level/region"
)

// ChunkLoadWorker loads and unloads chunks for the session
type ChunkLoadWorker struct {
	s        *StandardSession
	requests chan [3]int32 //x,z,radius
}

func NewChunkLoadWorker(s *StandardSession) *ChunkLoadWorker {
	return &ChunkLoadWorker{s: s, requests: make(chan [3]int32)}
}

func (c ChunkLoadWorker) SendChunksRadius(chunkX, chunkZ, radius int32) {
	c.requests <- [3]int32{chunkX, chunkZ, radius}
}

func (c ChunkLoadWorker) sendChunksRadius(chunkX, chunkZ, rad int32) error {
	c.s.load_ch_mu.Lock()
	defer c.s.load_ch_mu.Unlock()

	if rad == 0 {
		rad = c.s.ViewDistance()
	}

	var chunks int32

	if err := c.s.WritePacket(&play.ChunkBatchStart{}); err != nil {
		return err
	}

	for x := chunkX - rad; x < chunkX+rad; x++ {
		for z := chunkZ - rad; z < chunkZ+rad; z++ {
			if _, ok := c.s.loadedChunks[region.ChunkHash(x, z)]; ok {
				continue
			}
			c.s.loadedChunks[region.ChunkHash(x, z)] = struct{}{}
			chunk, err := c.s.Dimension().GetChunk(x, z)
			if err != nil {
				continue
			}

			if err := c.s.WritePacket(chunk.Encode(c.s.registryIndexes["minecraft:worldgen/biome"])); err != nil {
				return err
			}
			chunks++
		}
	}

	if err := c.s.WritePacket(&play.ChunkBatchFinished{
		BatchSize: chunks,
	}); err != nil {
		return err
	}
	c.s.awaitingChunkBatchAcknowledgement.Store(true)

	return nil
}

func (c ChunkLoadWorker) start() {
	go func() {
		for req := range c.requests {
			c.sendChunksRadius(req[0], req[1], req[2])
		}
	}()
}

func (session *StandardSession) sendSpawnChunks() error {
	viewDistance := session.ViewDistance()

	x, _, z := session.player.Position()
	chunkX, chunkZ := int32(x)>>4, int32(z)>>4

	if err := session.WritePacket(&play.SetCenterChunk{ChunkX: chunkX, ChunkZ: chunkZ}); err != nil {
		return err
	}

	var chunks int32

	if err := session.WritePacket(&play.ChunkBatchStart{}); err != nil {
		return err
	}

	for x := chunkX - viewDistance; x < chunkX+viewDistance; x++ {
		for z := chunkZ - viewDistance; z < chunkZ+viewDistance; z++ {
			c, err := session.Dimension().GetChunk(x, z)
			if err != nil {
				continue
			}

			if err := session.WritePacket(c.Encode(session.registryIndexes["minecraft:worldgen/biome"])); err != nil {
				return err
			}
			chunks++
		}
	}

	if err := session.WritePacket(&play.ChunkBatchFinished{
		BatchSize: chunks,
	}); err != nil {
		return err
	}

	session.awaitingChunkBatchAcknowledgement.Store(true)
	session.ChunkLoadWorker.start()

	return nil
}
