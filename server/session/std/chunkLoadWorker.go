package std

import "github.com/zeppelinmc/zeppelin/server/world/level/region"

// ChunkLoadWorker loads and unloads chunks for the session
type ChunkLoadWorker struct {
	s        *StandardSession
	requests chan [2]int32
}

func NewChunkLoadWorker(s *StandardSession) *ChunkLoadWorker {
	return &ChunkLoadWorker{s: s, requests: make(chan [2]int32)}
}

func (c ChunkLoadWorker) SendChunksRadius(chunkX, chunkZ int32) {
	c.requests <- [2]int32{chunkX, chunkZ}
}

func (c ChunkLoadWorker) sendChunksRadius(chunkX, chunkZ int32) error {
	viewDistance := c.s.ViewDistance()

	c.s.load_ch_mu.Lock()
	defer c.s.load_ch_mu.Unlock()

	for x := chunkX - viewDistance; x < chunkX+viewDistance; x++ {
		for z := chunkZ - viewDistance; z < chunkZ+viewDistance; z++ {
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
		}
	}

	return nil
}

func (c ChunkLoadWorker) start() {
	go func() {
		for pos := range c.requests {
			c.sendChunksRadius(pos[0], pos[1])
		}
	}()
}
