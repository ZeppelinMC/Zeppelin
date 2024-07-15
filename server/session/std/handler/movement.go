package handler

import (
	"math"

	"github.com/dynamitemc/aether/log"
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
		oldX, oldY, oldZ := s.Player().Position()
		oldChunkPosX, oldChunkPosZ := chunkPos(oldX, oldZ)
		newChunkPosX, newChunkPosZ := chunkPos(pk.X, pk.Z)

		if oldChunkPosX != newChunkPosX || oldChunkPosZ != newChunkPosZ {
			s.Conn().WritePacket(&play.SetCenterChunk{ChunkX: newChunkPosX, ChunkZ: newChunkPosZ})
		}

		distance := math.Sqrt((pk.X-oldX)*(pk.X-oldX) + (pk.Y-oldY)*(pk.Y-oldY) + (pk.Z-oldZ)*(pk.Z-oldZ))

		yaw, pitch := s.Player().Rotation()
		if distance > 100 {
			s.Teleport(oldX, oldY, oldZ, yaw, pitch)
			log.Infof("%s moved too quickly! (%f %f %f)\n", s.Username(), pk.X-oldX, pk.Y-oldY, pk.Z-oldZ)
			return
		}

		s.Broadcast().BroadcastPlayerMovement(s, pk.X, pk.Y, pk.Z, yaw, pitch)

		s.Player().SetPosition(pk.X, pk.Y, pk.Z)
	case *play.SetPlayerPositionAndRotation:
		oldX, oldY, oldZ := s.Player().Position()
		oldChunkPosX, oldChunkPosZ := chunkPos(oldX, oldZ)
		newChunkPosX, newChunkPosZ := chunkPos(pk.X, pk.Z)

		if oldChunkPosX != newChunkPosX || oldChunkPosZ != newChunkPosZ {
			s.Conn().WritePacket(&play.SetCenterChunk{ChunkX: newChunkPosX, ChunkZ: newChunkPosZ})
		}

		distance := math.Sqrt((pk.X-oldX)*(pk.X-oldX) + (pk.Y-oldY)*(pk.Y-oldY) + (pk.Z-oldZ)*(pk.Z-oldZ))

		if distance > 100 {
			s.Teleport(oldX, oldY, oldZ, pk.Yaw, pk.Pitch)
			log.Infof("%s moved too quickly! (%f %f %f)\n", s.Username(), pk.X-oldX, pk.Y-oldY, pk.Z-oldZ)
			return
		}

		s.Broadcast().BroadcastPlayerMovement(s, pk.X, pk.Y, pk.Z, pk.Yaw, pk.Pitch)

		s.Player().SetPosition(pk.X, pk.Y, pk.Z)
		s.Player().SetRotation(pk.Yaw, pk.Pitch)
	case *play.SetPlayerRotation:
		// you can never rotate too much :)
		x, y, z := s.Player().Position()

		s.Broadcast().BroadcastPlayerMovement(s, x, y, z, pk.Yaw, pk.Pitch)

		s.Player().SetRotation(pk.Yaw, pk.Pitch)
	}
}
