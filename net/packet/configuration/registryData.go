package configuration

import (
	"aether/net/io"
	"fmt"
	"reflect"
)

type RegistryData struct {
	RegistryId string
	Registry   any
}

func (RegistryData) ID() int32 {
	return 0x07
}

func (r *RegistryData) Encode(w io.Writer) error {
	if err := w.Identifier(r.RegistryId); err != nil {
		return err
	}
	fmt.Println(r.RegistryId)
	reg := reflect.ValueOf(r.Registry)
	if err := w.VarInt(int32(reg.Len())); err != nil {
		return err
	}
	for _, key := range reg.MapKeys() {
		v := reg.MapIndex(key)

		fmt.Println(key)
		if err := w.Identifier(key.String()); err != nil {
			return err
		}
		if err := w.Bool(!v.IsZero()); err != nil {
			return err
		}
		if !v.IsZero() {
			if err := w.NBT(v); err != nil {
				return err
			}
		}
	}
	return nil
}

func (d *RegistryData) Decode(r io.Reader) error {
	return nil
	//TODO
}
