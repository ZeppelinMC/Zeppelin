package play

import "github.com/zeppelinmc/zeppelin/net/io"

//clientbound
const PacketIdUpdateTags = 0x78

type UpdateTags struct {
	Tags map[string]map[string][]int32
}

func (UpdateTags) ID() int32 {
	return PacketIdUpdateTags
}

func (u *UpdateTags) Encode(w io.Writer) error {
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

func (u *UpdateTags) Decode(r io.Reader) error {
	return nil //TODO
}
