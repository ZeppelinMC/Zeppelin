package configuration

import "aether/net/io"

type Entry struct {
	ID   string
	Data []byte
}

type RegistryData struct {
	EntryID string
	Entries []Entry
}

func (RegistryData) ID() int32 {
	return 0x07
}

func (r *RegistryData) Encode(w io.Writer) error {
	if err := w.Identifier(r.EntryID); err != nil {
		return err
	}
	if err := w.VarInt(int32(len(r.Entries))); err != nil {
		return err
	}
	for _, entry := range r.Entries {
		if err := w.Identifier(entry.ID); err != nil {
			return err
		}
		if err := w.Bool(len(entry.Data) != 0); err != nil {
			return err
		}
		if len(entry.Data) != 0 {
			//TODO
		}
	}
	return nil
}

func (d *RegistryData) Decode(r io.Reader) error {
	if err := r.Identifier(&d.EntryID); err != nil {
		return err
	}
	var length int32
	if _, err := r.VarInt(&length); err != nil {
		return err
	}
	d.Entries = make([]Entry, length)
	for _, entry := range d.Entries {
		if err := r.Identifier(&entry.ID); err != nil {
			return err
		}
		var hasData bool
		if err := r.Bool(&hasData); err != nil {
			return err
		}
		if hasData {
			//TODO
		}
	}
	return nil
}
