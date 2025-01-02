package slot

import (
	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
	"github.com/zeppelinmc/zeppelin/protocol/text"
)

func decode(r encoding.Reader, comp *Component) error {
	switch comp.Type {
	case CustomData:
		comp.Data = make(map[string]any)
		if err := r.NBT(&comp.Data); err != nil {
			return err
		}
	case MaxStackSize, Damage, MaxDamage, Rarity:
		var data int32
		if _, err := r.VarInt(&data); err != nil {
			return err
		}
		comp.Data = data
	case Unbreakable:
		var data bool
		if err := r.Bool(&data); err != nil {
			return err
		}
		comp.Data = data
	case ItemName, CustomName:
		var data text.TextComponent
		if err := r.TextComponent(&data); err != nil {
			return err
		}
		comp.Data = data
	case Lore:
		var number int32
		if _, err := r.VarInt(&number); err != nil {
			return err
		}
		var data = make([]text.TextComponent, number)
		for _, line := range data {
			if err := r.TextComponent(&line); err != nil {
				return err
			}
		}
		comp.Data = data
	}

	return nil
}
