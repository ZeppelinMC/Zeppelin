package qnbt

import "fmt"

func (d *Decoder) decodeMapString(m map[string]string) error {
	var name string
	for {
		tag, err := d.rd.readBytesString(1)
		if err != nil {
			return err
		}

		if tag == end_t_b {
			return nil
		}
		if err := d.readStringNonCopy(&name); err != nil {
			return err
		}

		if tag != string_t_b {
			return fmt.Errorf("expected string tag, got %s", tagName(tag))
		}

		var v string
		if err := d.readString(&v); err != nil {
			return err
		}
		m[name] = v
	}
}
