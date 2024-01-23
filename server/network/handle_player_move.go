package network

import (
	"github.com/aimjel/minecraft/protocol"
	"github.com/aimjel/minecraft/protocol/metadata"
	"math"

	"github.com/aimjel/minecraft/packet"
)

func DegreesToAngle(degrees float32) byte {
	return byte(math.Round(float64(degrees) * (256.0 / 360.0)))
}

func PositionIsValid(x, y, z float64) bool {
	return !math.IsNaN(x) && !math.IsNaN(y) && !math.IsNaN(z) &&
		!math.IsInf(x, 0) && !math.IsInf(y, 0) && !math.IsInf(z, 0)
}

// delta calculates delta between 3d coordinates.
func delta(currX, currY, currZ, prevX, prevY, prevZ float64) (dx int16, dy int16, dz int16) {
	dx = int16(((currX * 32) - prevX*32) * 128)
	dy = int16(((currY * 32) - prevY*32) * 128)
	dz = int16(((currZ * 32) - prevZ*32) * 128)
	return
}

func HandlePlayerMovement(s *Session, p packet.Packet) {
	switch pk := p.(type) {

	case *packet.PlayerPosition:
		prevX, prevY, prevZ := s.state.Position()
		yaw, pitch := s.state.Rotation()

		s.state.Move(pk.X, pk.FeetY, pk.Z)

		dx, dy, dz := delta(pk.X, pk.FeetY, pk.Z, prevX, prevY, prevZ)

		yawAngle, pitchAngle := DegreesToAngle(yaw), DegreesToAngle(pitch)

		s.b.BroadcastMovement(s, dx, dy, dz, yawAngle, pitchAngle, sendPosition)

		s.handleChunkBorder(prevX, prevZ, pk.X, pk.Z)

	case *packet.PlayerPositionRotation:
		prevX, prevY, prevZ := s.state.Position()

		s.state.Move(pk.X, pk.FeetY, pk.Z)
		s.state.Rotate(pk.Yaw, pk.Pitch)

		dx, dy, dz := delta(pk.X, pk.FeetY, pk.Z, prevX, prevY, prevZ)

		yawAngle, pitchAngle := DegreesToAngle(pk.Yaw), DegreesToAngle(pk.Pitch)

		s.b.BroadcastMovement(s, dx, dy, dz, yawAngle, pitchAngle, sendPosition|sendRotation)

		s.handleChunkBorder(prevX, prevZ, pk.X, pk.Z)

	case *packet.PlayerRotation:
		s.state.Rotate(pk.Yaw, pk.Pitch)

		yawAngle, pitchAngle := DegreesToAngle(pk.Yaw), DegreesToAngle(pk.Pitch)

		s.b.BroadcastMovement(s, 0, 0, 0, yawAngle, pitchAngle, sendRotation)
	}
}

func createRange(p1, p2, vd int32, r1, r2 *[2]int32) {
	if p1 != p2 {
		if p2 > p1 {
			r1[0] += vd
			r1[1] += vd
		} else {
			r1[0] -= vd
			r1[1] -= vd
		}

		r2[0] -= vd
		r2[1] += vd
	}
}

func (s *Session) handleChunkBorder(x1, z1, x2, z2 float64) {
	prevX, prevZ := int32(x1/16), int32(z1/16)
	currX, currZ := int32(x2/16), int32(z2/16)

	if prevX == currX && prevZ == currZ {
		return
	}

	//holds a range of x and z coordinates to load for the client
	var x, z = [2]int32{prevX, prevX}, [2]int32{prevZ, prevZ}
	createRange(prevX, currX, s.ViewDistance, &x, &z)
	createRange(prevZ, currZ, s.ViewDistance, &z, &x)

	s.conn.SendPacket(&packet.SetCenterChunk{
		ChunkX: currX,
		ChunkZ: currZ})

	//used for encoding height-map, sections and light info
	data := protocol.GetBuffer(1024 * 21)
	defer protocol.PutBuffer(data)

	for loadX := x[0]; loadX <= x[1]; loadX++ {
		for loadZ := z[0]; loadZ <= z[1]; loadZ++ {
			c, err := s.state.Dimension().Chunk(loadX, loadZ)
			if err != nil {
				panic(err)
			}

			c.NetworkEncode(data)

			s.conn.SendPacket(&packet.ChunkData{
				X:    c.X,
				Z:    c.Z,
				Data: data.Bytes(),
			})

			data.Reset()
		}
	}
}

func HandlePlayerCommand(s *Session, p packet.Packet) {
	if pk, ok := p.(*packet.PlayerCommandServer); ok {
		switch pk.ActionID {
		case 0x00: // start sneaking
			s.b.BroadcastEntityMetaData(s, (metadata.Entity{}).Crouch(true))
		case 0x01: // stop sneaking
			s.b.BroadcastEntityMetaData(s, (metadata.Entity{}).Crouch(false))
		case 0x03: // start sprinting
			// the client sends this also when it's swimming, so we must check if it is
			var swimming bool
			if !swimming {
				s.b.BroadcastEntityMetaData(s, (metadata.Entity{}).Sprinting(true))
			}

		case 0x04: // stop sprinting
			s.b.BroadcastEntityMetaData(s, (metadata.Entity{}).Sprinting(false))
		}
	}
}
