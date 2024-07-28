package play

import (
	"github.com/zeppelinmc/zeppelin/net/io"
	"github.com/zeppelinmc/zeppelin/net/metadata"
	"github.com/zeppelinmc/zeppelin/text"
)

// clientbound
const PacketIdSetEntityMetadata = 0x58

type SetEntityMetadata struct {
	EntityId int32
	Metadata metadata.Metadata
}

func (SetEntityMetadata) ID() int32 {
	return PacketIdSetEntityMetadata
}

func (s *SetEntityMetadata) Encode(w io.Writer) error {
	if err := w.VarInt(s.EntityId); err != nil {
		return err
	}
	for index, value := range s.Metadata {
		if err := w.Ubyte(index); err != nil {
			return err
		}
		switch val := value.(type) {
		case metadata.Byte:
			if err := w.VarInt(0); err != nil {
				return err
			}
			if err := w.Byte(int8(val)); err != nil {
				return err
			}
		case metadata.VarInt:
			if err := w.VarInt(1); err != nil {
				return err
			}
			if err := w.VarInt(int32(val)); err != nil {
				return err
			}
		case metadata.VarLong:
			if err := w.VarInt(2); err != nil {
				return err
			}
			if err := w.VarLong(int64(val)); err != nil {
				return err
			}
		case metadata.Float:
			if err := w.VarInt(3); err != nil {
				return err
			}
			if err := w.Float(float32(val)); err != nil {
				return err
			}
		case metadata.String:
			if err := w.VarInt(4); err != nil {
				return err
			}
			if err := w.String(string(val)); err != nil {
				return err
			}
		case metadata.TextComponent:
			if err := w.VarInt(5); err != nil {
				return err
			}
			if err := w.TextComponent(text.TextComponent(val)); err != nil {
				return err
			}
		case metadata.OptionalTextComponent:
			if err := w.VarInt(6); err != nil {
				return err
			}
			if err := w.Bool(val != nil); err != nil {
				return err
			}
			if val != nil {
				if err := w.TextComponent(text.TextComponent(*val)); err != nil {
					return err
				}
			}
		case metadata.Slot:
			if err := w.VarInt(7); err != nil {
				return err
			}
			if err := val.Encode(w); err != nil {
				return err
			}
		case metadata.Boolean:
			if err := w.VarInt(8); err != nil {
				return err
			}
			if err := w.Bool(bool(val)); err != nil {
				return err
			}
		case metadata.Rotations:
			if err := w.VarInt(9); err != nil {
				return err
			}
			if err := w.Float(float32(val[0])); err != nil {
				return err
			}
			if err := w.Float(float32(val[1])); err != nil {
				return err
			}
			if err := w.Float(float32(val[2])); err != nil {
				return err
			}
		case metadata.Position:
			if err := w.VarInt(10); err != nil {
				return err
			}
			if err := w.Position(val[0], val[1], val[2]); err != nil {
				return err
			}
		case metadata.OptionalPosition:
			if err := w.VarInt(11); err != nil {
				return err
			}
			if err := w.Bool(val != nil); err != nil {
				return err
			}
			if val != nil {
				if err := w.Position(val[0], val[1], val[2]); err != nil {
					return err
				}
			}
		case metadata.Direction:
			if err := w.VarInt(12); err != nil {
				return err
			}
			if err := w.VarInt(int32(val)); err != nil {
				return err
			}
		case metadata.OptionalUUID:
			if err := w.VarInt(13); err != nil {
				return err
			}
			if err := w.Bool(val != nil); err != nil {
				return err
			}
			if val != nil {
				if err := w.UUID(*val); err != nil {
					return err
				}
			}
		case metadata.BlockState:
			if err := w.VarInt(14); err != nil {
				return err
			}
			if err := w.VarInt(int32(val)); err != nil {
				return err
			}
		case metadata.OptionalBlockState:
			if err := w.VarInt(15); err != nil {
				return err
			}
			if err := w.VarInt(int32(val)); err != nil {
				return err
			}
		case metadata.VillagerData:
			if err := w.VarInt(18); err != nil {
				return err
			}
			if err := w.Float(float32(val[0])); err != nil {
				return err
			}
			if err := w.Float(float32(val[1])); err != nil {
				return err
			}
			if err := w.Float(float32(val[2])); err != nil {
				return err
			}
		case metadata.OptionalVarInt:
			if err := w.VarInt(19); err != nil {
				return err
			}
			if err := w.VarInt(int32(val)); err != nil {
				return err
			}
		case metadata.Pose:
			if err := w.VarInt(21); err != nil {
				return err
			}
			if err := w.VarInt(int32(val)); err != nil {
				return err
			}
		case metadata.CatVariant:
			if err := w.VarInt(21); err != nil {
				return err
			}
			if err := w.VarInt(int32(val)); err != nil {
				return err
			}
		case metadata.FrogVariant:
			if err := w.VarInt(22); err != nil {
				return err
			}
			if err := w.VarInt(int32(val)); err != nil {
				return err
			}
		case metadata.OptionalGlobalPosition:
			if err := w.VarInt(23); err != nil {
				return err
			}
			if err := w.Bool(val != nil); err != nil {
				return err
			}
			if val != nil {
				if err := w.Identifier(string(val.DimensionIdentifier)); err != nil {
					return err
				}
				if err := w.Position(val.Position[0], val.Position[1], val.Position[2]); err != nil {
					return err
				}
			}
		case metadata.PaintingVariant:
			if err := w.VarInt(24); err != nil {
				return err
			}
			if err := w.VarInt(int32(val)); err != nil {
				return err
			}
		case metadata.SnifferState:
			if err := w.VarInt(25); err != nil {
				return err
			}
			if err := w.VarInt(int32(val)); err != nil {
				return err
			}
		case metadata.Vector3:
			if err := w.VarInt(26); err != nil {
				return err
			}
			if err := w.Float(float32(val[0])); err != nil {
				return err
			}
			if err := w.Float(float32(val[1])); err != nil {
				return err
			}
			if err := w.Float(float32(val[2])); err != nil {
				return err
			}
		case metadata.Quatermion:
			if err := w.VarInt(27); err != nil {
				return err
			}
			if err := w.Float(float32(val[0])); err != nil {
				return err
			}
			if err := w.Float(float32(val[1])); err != nil {
				return err
			}
			if err := w.Float(float32(val[2])); err != nil {
				return err
			}
			if err := w.Float(float32(val[3])); err != nil {
				return err
			}
		case metadata.NBT:
			if err := w.VarInt(16); err != nil {
				return err
			}
			if err := w.NBT(val); err != nil {
				return err
			}
		default:
			continue
		}
	}
	return w.Ubyte(0xFF)
}

func (*SetEntityMetadata) Decode(r io.Reader) error {
	return nil //TODO
}
