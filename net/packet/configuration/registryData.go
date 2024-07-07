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

	reg := reflect.ValueOf(r.Registry)

	if err := w.VarInt(int32(reg.Len())); err != nil {
		return err
	}
	for _, key := range reg.MapKeys() {
		v := reg.MapIndex(key).Interface()

		if err := w.Identifier(key.String()); err != nil {
			return err
		}

		fmt.Println("------", key.String(), "------")

		if err := w.Bool(true); err != nil {
			return err
		}

		if err := w.NBT(v); err != nil {
			return err
		}

		//break

		/*if key.String() == "minecraft:chat" {

			var buf = bytes.NewBuffer(nil)
			enc := nbt.NewEncoder(buf)
			enc.WriteRootName(false)
			enc.Encode("", v)

			fmt.Print("[")
			for i, k := range buf.Bytes() {
				if i != 0 {
					fmt.Print(" ")
				}
				fmt.Printf("%02x", k)
			}
			fmt.Println("]")

		}*/
	}
	return nil
}

func (d *RegistryData) Decode(r io.Reader) error {
	return nil
	//TODO
}
