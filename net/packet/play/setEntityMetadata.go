package play

import (
	"fmt"
	"reflect"

	"github.com/dynamitemc/aether/chat"
	"github.com/dynamitemc/aether/net/io"
	"github.com/google/uuid"
)

// clientbound
const PacketIdSetEntityMetadata = 0x58

type SetEntityMetadata struct {
	EntityId int32
	// index -> value
	Metadata map[byte]any
}

type (
	Byte    int8    //0
	VarInt  int32   //1
	VarLong int64   //2
	Float   float32 //3
	String  string  //4

	TextComponent         chat.TextComponent  //5
	OptionalTextComponent *chat.TextComponent //6
	//Slot //7
	Boolean            bool       //8
	Rotations          [3]Float   //9
	Position           [3]int32   //10
	OptionalPosition   *[3]int32  //11
	Direction          VarInt     //12
	OptionalUUID       *uuid.UUID //13
	BlockState         VarInt     //14
	OptionalBlockState VarInt     //15
	NBT                any        //16
	// Particle //17
	VillagerData   [3]Float //18 | [type, profession, level]
	OptionalVarInt VarInt   //19
	Pose           VarInt   //20
	CatVariant     VarInt   //21
	FrogVariant    VarInt   //22

	GlobalPosition struct {
		DimensionIdentifier String
		Position            Position
	}
	OptionalGlobalPosition *GlobalPosition //23

	PaintingVariant VarInt   //24
	SnifferState    VarInt   //25
	Vector3         [3]Float //26
	Quatermion      [4]Float //27
)

const (
	Standing Pose = iota
	FallFlying
	Sleeping
	Swimming
	SpinAttack
	Sneaking
	LongJumping
	Dying
	Croaking
	UsingTongue
	Sitting
	Roaring
	Sniffing
	Emerging
	Digging
)

const (
	SnifferIdling SnifferState = iota
	SnifferFeelingHappy
	SnifferScenting
	SnifferSniffing
	SnifferSearching
	SnifferDigging
	SnifferRising
)

const (
	VillagerTypeDesert = iota
	VillagerTypeJungle
	VillagerTypePlains
	VillagerTypeSavanna
	VillagerTypeSnow
	VillagerTypeSwamp
	VillagerTypeTaiga
)

const (
	VillagerProfessionNone = iota
	VillagerProfessionArmorer
	VillagerProfessionButcher
	VillagerProfessionCartographer
	VillagerProfessionCleric
	VillagerProfessionFarmer
	VillagerProfessionFisherman
	VillagerProfessionFletcher
	VillagerProfessionLeatherworker
	VillagerProfessionLibrarian
	VillagerProfessionMason // hi mason
	VillagerProfessionNitwit
	VillagerProfessionShepherd
	VillagerProfessionToolsmith
	VillagerProfessionWeaponsmith
)

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
		fmt.Println(index, reflect.TypeOf(value), value)
		switch val := value.(type) {
		case Byte:
			if err := w.VarInt(0); err != nil {
				return err
			}
			if err := w.Byte(int8(val)); err != nil {
				return err
			}
		case VarInt:
			if err := w.VarInt(1); err != nil {
				return err
			}
			if err := w.VarInt(int32(val)); err != nil {
				return err
			}
		case VarLong:
			if err := w.VarInt(2); err != nil {
				return err
			}
			if err := w.VarLong(int64(val)); err != nil {
				return err
			}
		case Float:
			if err := w.VarInt(3); err != nil {
				return err
			}
			if err := w.Float(float32(val)); err != nil {
				return err
			}
		case String:
			if err := w.VarInt(4); err != nil {
				return err
			}
			if err := w.String(string(val)); err != nil {
				return err
			}
		case TextComponent:
			if err := w.VarInt(5); err != nil {
				return err
			}
			if err := w.TextComponent(chat.TextComponent(val)); err != nil {
				return err
			}
		case OptionalTextComponent:
			if err := w.VarInt(6); err != nil {
				return err
			}
			if err := w.Bool(val != nil); err != nil {
				return err
			}
			if val != nil {
				if err := w.TextComponent(chat.TextComponent(*val)); err != nil {
					return err
				}
			}
		//case Slot:
		case Boolean:
			if err := w.VarInt(8); err != nil {
				return err
			}
			if err := w.Bool(bool(val)); err != nil {
				return err
			}
		case Rotations:
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
		case Position:
			if err := w.VarInt(10); err != nil {
				return err
			}
			if err := w.Position(val[0], val[1], val[2]); err != nil {
				return err
			}
		case OptionalPosition:
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
		case Direction:
			if err := w.VarInt(12); err != nil {
				return err
			}
			if err := w.VarInt(int32(val)); err != nil {
				return err
			}
		case OptionalUUID:
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
		case BlockState:
			if err := w.VarInt(14); err != nil {
				return err
			}
			if err := w.VarInt(int32(val)); err != nil {
				return err
			}
		case OptionalBlockState:
			if err := w.VarInt(15); err != nil {
				return err
			}
			if err := w.VarInt(int32(val)); err != nil {
				return err
			}
		case VillagerData:
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
		case OptionalVarInt:
			if err := w.VarInt(19); err != nil {
				return err
			}
			if err := w.VarInt(int32(val)); err != nil {
				return err
			}
		case Pose:
			if err := w.VarInt(21); err != nil {
				return err
			}
			if err := w.VarInt(int32(val)); err != nil {
				return err
			}
		case CatVariant:
			if err := w.VarInt(21); err != nil {
				return err
			}
			if err := w.VarInt(int32(val)); err != nil {
				return err
			}
		case FrogVariant:
			if err := w.VarInt(22); err != nil {
				return err
			}
			if err := w.VarInt(int32(val)); err != nil {
				return err
			}
		case OptionalGlobalPosition:
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
		case PaintingVariant:
			if err := w.VarInt(24); err != nil {
				return err
			}
			if err := w.VarInt(int32(val)); err != nil {
				return err
			}
		case SnifferState:
			if err := w.VarInt(25); err != nil {
				return err
			}
			if err := w.VarInt(int32(val)); err != nil {
				return err
			}
		case Vector3:
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
		case Quatermion:
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
		case NBT:
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
