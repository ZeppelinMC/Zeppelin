package handler

import (
	"math"

	"github.com/dynamitemc/aether/net"
	"github.com/dynamitemc/aether/net/packet"
	"github.com/dynamitemc/aether/net/packet/play"
	"github.com/dynamitemc/aether/server/session"
)

func init() {
	session.Handlers[[2]int32{net.PlayState, play.PacketIdSetPlayerPosition}] = handleMovement
	session.Handlers[[2]int32{net.PlayState, play.PacketIdSetPlayerPositionAndRotation}] = handleMovement
	session.Handlers[[2]int32{net.PlayState, play.PacketIdSetPlayerRotation}] = handleMovement
}

func chunkPos(pos float64) int32 {
	return int32(math.Floor(pos / 16))
}

func handleMovement(s *session.Session, p packet.Packet) {
	switch pk := p.(type) {
	case *play.SetPlayerPosition:
		oldX, _, oldZ := s.Player.Position()
		oldChunkX, oldChunkZ := chunkPos(oldX), chunkPos(oldZ)
		chunkX, chunkZ := chunkPos(pk.X), chunkPos(pk.Z)

		s.Player.SetPosition(pk.X, pk.Y, pk.Z)
		if oldChunkX != chunkX || oldChunkZ != chunkZ {
			s.Conn.WritePacket(&play.SetCenterChunk{ChunkX: chunkX, ChunkZ: chunkZ})
		}
	case *play.SetPlayerPositionAndRotation:
		oldX, _, oldZ := s.Player.Position()
		oldChunkX, oldChunkZ := chunkPos(oldX), chunkPos(oldZ)
		chunkX, chunkZ := chunkPos(pk.X), chunkPos(pk.Z)

		s.Player.SetPosition(pk.X, pk.Y, pk.Z)
		s.Player.SetRotation(pk.Yaw, pk.Pitch)
		if oldChunkX != chunkX || oldChunkZ != chunkZ {
			s.Conn.WritePacket(&play.SetCenterChunk{ChunkX: chunkX, ChunkZ: chunkZ})
		}
	case *play.SetPlayerRotation:
		s.Player.SetRotation(pk.Yaw, pk.Pitch)
	}
}
