package configuration

import (
	"aether/net/io"
	"aether/net/registry"
	"reflect"
)

func ConstructRegistryPackets() []*RegistryData {
	regDatas := make([]*RegistryData, 0, len(registry.RegistryMap))
	for key, registry := range registry.RegistryMap {
		regDatas = append(regDatas, &RegistryData{
			RegistryId: key,
			Registry:   registry,
		})
	}
	return regDatas
}

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

	reg := reflect.ValueOf(r.Registry)

	var i int32
	var m map[string]int32
	if r.RegistryId == "minecraft:worldgen/biome" {
		m = make(map[string]int32)
	}
	switch reg.Kind() {
	case reflect.Map:
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

			if r.RegistryId == "minecraft:worldgen/biome" {
				m[key.String()] = i
			}

			i++
		}
		if r.RegistryId == "minecraft:worldgen/biome" {
			registry.BiomeId.SetMap(m)
		}
	case reflect.Struct:
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
		}
	}
	return nil
}

func (d *RegistryData) Decode(r io.Reader) error {
	return nil
	//TODO
}
