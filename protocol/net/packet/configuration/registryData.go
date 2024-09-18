package configuration

import (
	"reflect"
	"sync"

	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
	"github.com/zeppelinmc/zeppelin/protocol/net/registry"
)

var RegistryPackets = make([]*RegistryData, 0, len(registry.Registries))

func init() {
	for key, registry := range registry.Registries {
		RegistryPackets = append(RegistryPackets, &RegistryData{
			RegistryId: key,
			Registry:   registry,
		})
	}
}

var RegistryPacketsMutex sync.Mutex

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

func (r *RegistryData) Encode(w encoding.Writer) error {
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

func (d *RegistryData) Decode(r encoding.Reader) error {
	return nil
	//TODO
}
