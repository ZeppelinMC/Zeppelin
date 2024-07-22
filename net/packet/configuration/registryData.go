package configuration

import (
	"reflect"

	"github.com/zeppelinmc/zeppelin/net/io"
	"github.com/zeppelinmc/zeppelin/net/registry"
)

var RegistryPackets = make([]*RegistryData, 0, len(registry.RegistryMap))

func init() {
	for key, registry := range registry.RegistryMap {
		RegistryPackets = append(RegistryPackets, &RegistryData{
			RegistryId: key,
			Registry:   registry,
		})
	}
}

// clientbound
const PacketIdRegistryData = 0x07

type RegistryData struct {
	RegistryId string
	Registry   any

	Indexes []string
}

func (RegistryData) ID() int32 {
	return 0x07
}

func (r *RegistryData) Encode(w io.Writer) error {
	if err := w.Identifier(r.RegistryId); err != nil {
		return err
	}

	reg := reflect.ValueOf(r.Registry)

	switch reg.Kind() {
	case reflect.Map:
		r.Indexes = make([]string, 0, reg.Len())
		if err := w.VarInt(int32(reg.Len())); err != nil {
			return err
		}
		for _, key := range reg.MapKeys() {
			v := reg.MapIndex(key).Interface()

			if err := w.Identifier(key.String()); err != nil {
				return err
			}
			if err := w.Bool(true); err != nil {
				return err
			}

			if err := w.NBT(v); err != nil {
				return err
			}
			r.Indexes = append(r.Indexes, key.String())
		}
	case reflect.Struct:
		r.Indexes = make([]string, 0, reg.NumField())

		if err := w.VarInt(int32(reg.NumField())); err != nil {
			return err
		}
		for i := 0; i < reg.NumField(); i++ {
			tf := reg.Type().Field(i)
			v := reg.Field(i).Interface()
			nbtname := tf.Tag.Get("nbt")

			if err := w.Identifier(nbtname); err != nil {
				return err
			}
			if err := w.Bool(true); err != nil {
				return err
			}

			if err := w.NBT(v); err != nil {
				return err
			}
			r.Indexes = append(r.Indexes, nbtname)
		}
	}
	return nil
}

func (d *RegistryData) Decode(r io.Reader) error {
	return nil
	//TODO
}
