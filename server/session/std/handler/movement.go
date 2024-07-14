package handler

import (
	"math"

	"github.com/dynamitemc/aether/net"
	"github.com/dynamitemc/aether/net/packet"
	"github.com/dynamitemc/aether/net/packet/play"
	"github.com/dynamitemc/aether/server/session/std"
)

func init() {
	std.RegisterHandler(net.PlayState, play.PacketIdSetPlayerPosition, handleMovement)
	std.RegisterHandler(net.PlayState, play.PacketIdSetPlayerPositionAndRotation, handleMovement)
	std.RegisterHandler(net.PlayState, play.PacketIdSetPlayerRotation, handleMovement)
}

func chunkPos(x, z float64) (cx, cz int32) {
	return int32(math.Floor(x / 16)), int32(math.Floor(z / 16))
}

func handleMovement(s *std.StandardSession, p packet.Packet) {
	switch pk := p.(type) {
	case *play.SetPlayerPosition:
		oldX, _, oldZ := s.Player().Position()
		oldChunkPosX, oldChunkPosZ := chunkPos(oldX, oldZ)
		newChunkPosX, newChunkPosZ := chunkPos(pk.X, pk.Z)

		if oldChunkPosX != newChunkPosX || oldChunkPosZ != newChunkPosZ {
			s.Conn().WritePacket(&play.SetCenterChunk{ChunkX: newChunkPosX, ChunkZ: newChunkPosZ})
		}

		s.Player().SetPosition(pk.X, pk.Y, pk.Z)
	case *play.SetPlayerPositionAndRotation:
		oldX, _, oldZ := s.Player().Position()
		oldChunkPosX, oldChunkPosZ := chunkPos(oldX, oldZ)
		newChunkPosX, newChunkPosZ := chunkPos(pk.X, pk.Z)

		if oldChunkPosX != newChunkPosX || oldChunkPosZ != newChunkPosZ {
			s.Conn().WritePacket(&play.SetCenterChunk{ChunkX: newChunkPosX, ChunkZ: newChunkPosZ})
		}

		s.Player().SetPosition(pk.X, pk.Y, pk.Z)
		s.Player().SetRotation(pk.Yaw, pk.Pitch)
	case *play.SetPlayerRotation:
		s.Player().SetRotation(pk.Yaw, pk.Pitch)
	}
}
