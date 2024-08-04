package handler

import (
	"math"

	"github.com/zeppelinmc/zeppelin/net"
	"github.com/zeppelinmc/zeppelin/net/metadata"
	"github.com/zeppelinmc/zeppelin/net/packet"
	"github.com/zeppelinmc/zeppelin/net/packet/play"
	"github.com/zeppelinmc/zeppelin/server/item"
	"github.com/zeppelinmc/zeppelin/server/session/std"
	"github.com/zeppelinmc/zeppelin/server/world/block"
	"github.com/zeppelinmc/zeppelin/server/world/block/pos"
	"github.com/zeppelinmc/zeppelin/util"
)

func init() {
	std.RegisterHandler(net.PlayState, play.PacketIdUseItemOn, handleUseItemOn)
}

func handleUseItemOn(s *std.StandardSession, pk packet.Packet) {
	use, ok := pk.(*play.UseItemOn)
	if !ok {
		return
	}
	dimension := s.Dimension()

	x1, y1, z1 := s.Player().Position()
	x2, y2, z2 := float64(use.BlockX), float64(use.BlockY), float64(use.BlockZ)

	distance := math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2) + (z1-z2)*(z1-z2))
	if distance > 4 {
		return
	}

	b, err := dimension.Block(use.BlockX, use.BlockY, use.BlockZ)
	if err != nil {
		return
	}

	usable, ok := b.(block.Usable)
	if ok && s.Player().MetadataIndex(metadata.PoseIndex) != metadata.Sneaking {
		usable.Use(s, *use, dimension)
		return
	}

	itemInHand, ok := s.Player().Inventory().Slot(item.DataSlot(s.Player().SelectedItemSlot()))
	if !ok {
		return
	}
	blockSet, ok := itemInHand.Block()
	if !ok {
		return
	}

	var blockX, blockY, blockZ = use.BlockX, use.BlockY, use.BlockZ
	yaw, _ := s.Player().Rotation()

	var axis block.Axis
	var facing block.Direction

	switch use.Face {
	case play.FaceBottom:
		blockY--
		axis, facing = "y", util.YawDirection(yaw)
	case play.FaceTop:
		blockY++
		axis, facing = "y", util.YawDirection(yaw)
	case play.FaceNorth:
		blockZ--
		axis, facing = "z", "north"
	case play.FaceSouth:
		blockZ++
		axis, facing = "z", "south"
	case play.FaceWest:
		blockX--
		axis, facing = "x", "west"
	case play.FaceEast:
		blockX++
		axis, facing = "x", "east"
	}

	currentBlock, err := dimension.Block(blockX, blockY, blockZ)
	if err == nil {
		name, _ := currentBlock.Encode()
		if name != "minecraft:air" {
			return
		}
	}

	blockSet = blockSet.New(map[string]string{"axis": axis, "facing": facing})

	s.Conn().WritePacket(&play.AcknowledgeBlockChange{SequenceId: use.Sequence})
	pos := pos.New(blockX, blockY, blockZ)

	dimension.SetBlock(pos, blockSet, true)

	blockEntityHaver, ok := blockSet.(block.BlockEntityHaver)
	if !ok {
		return
	}
	dimension.SetBlockEntity(pos, blockEntityHaver.BlockEntity(pos))
}
