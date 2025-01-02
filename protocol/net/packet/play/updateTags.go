package play

import (
	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
)

// clientbound
const PacketIdUpdateTags = 0x78

type UpdateTags struct {
	Tags map[string]map[string][]int32
}

func (UpdateTags) ID() int32 {
	return PacketIdUpdateTags
}

func (u *UpdateTags) Encode(w encoding.Writer) error {
	if err := w.VarInt(int32(len(u.Tags))); err != nil {
		return err
	}
	for reg, tag := range u.Tags {
		if err := w.Identifier(reg); err != nil {
			return err
		}
		if err := w.VarInt(int32(len(tag))); err != nil {
			return err
		}
		for id, tag := range tag {
			if err := w.Identifier(id); err != nil {
				return err
			}
			if err := w.VarInt(int32(len(tag))); err != nil {
				return err
			}
			for _, entry := range tag {
				if err := w.VarInt(entry); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (u *UpdateTags) Decode(r encoding.Reader) error {
	var length int32
	if _, err := r.VarInt(&length); err != nil {
		return err
	}
	u.Tags = make(map[string]map[string][]int32, length)

	for i := int32(0); i < length; i++ {
		var registry string
		if err := r.String(&registry); err != nil {
			return err
		}
		var length int32
		if _, err := r.VarInt(&length); err != nil {
			return err
		}

		u.Tags[registry] = make(map[string][]int32, length)
		for i := int32(0); i < length; i++ {
			var tagName string
			if err := r.String(&tagName); err != nil {
				return err
			}
			var count int32
			if _, err := r.VarInt(&count); err != nil {
				return err
			}
			u.Tags[registry][tagName] = make([]int32, count)
			for i := int32(0); i < count; i++ {
				if _, err := r.VarInt(&u.Tags[registry][tagName][i]); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
