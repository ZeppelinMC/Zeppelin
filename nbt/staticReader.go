package nbt

import (
	"bytes"
	"fmt"
	"io"
	"unsafe"
)

/*
Reader is an NBT decoder that uses no reflection. it should be used for static structures.

The read functions read the type id, name and value
*/
type StaticReader struct {
	r *bytes.Reader
}

func NewStaticReader(r *bytes.Reader) StaticReader {
	return StaticReader{r}
}

func (r StaticReader) Skip(bytes int64) {
	r.r.Seek(bytes, io.SeekCurrent)
}

// reads a byte
func (r StaticReader) readByte() (byte, error) {
	var data [1]byte
	_, err := r.r.Read(data[:])

	return data[0], err
}

// reads a short, big endian
func (r StaticReader) readShort() (int16, error) {
	var data [2]byte
	_, err := r.r.Read(data[:])

	return int16(data[0])<<8 | int16(data[1]), err
}

// reads an int, big endian
func (r StaticReader) readInt() (int32, error) {
	var data [4]byte
	_, err := r.r.Read(data[:])

	return int32(data[0])<<24 | int32(data[1])<<16 | int32(data[2])<<8 | int32(data[3]), err
}

// reads a long, big endian
func (r StaticReader) readLong() (int64, error) {
	var data [8]byte
	_, err := r.r.Read(data[:])

	return int64(data[0])<<56 | int64(data[1])<<48 | int64(data[2])<<40 | int64(data[3])<<32 | int64(data[4])<<24 | int64(data[5])<<16 | int64(data[6])<<8 | int64(data[7]), err
}

// reads a string prefixed by a short of its length
func (r StaticReader) readString() (string, error) {
	length, err := r.readShort()
	if err != nil {
		return "", err
	}
	var data = make([]byte, length)

	if _, err := r.r.Read(data); err != nil {
		return "", err
	}

	return *(*string)(unsafe.Pointer(&data)), nil
}

func (r StaticReader) ReadRoot(withName bool) (typeId byte, name string, err error) {
	typeId, err = r.readByte()
	if err != nil {
		return typeId, "", err
	}
	if withName {
		name, err = r.readString()
	}

	return
}

func (r StaticReader) Byte(dst *byte) (string, error) {
	typeId, err := r.readByte()
	if err != nil {
		return "", err
	}
	if typeId != Byte {
		return "", fmt.Errorf("tried reading TAG_Byte, got %d", typeId)
	}

	name, err := r.readString()
	if err != nil {
		return "", err
	}

	*dst, err = r.readByte()
	return name, err
}

func (r StaticReader) Short(dst *int16) (string, error) {
	typeId, err := r.readByte()
	if err != nil {
		return "", err
	}
	if typeId != Short {
		return "", fmt.Errorf("tried reading TAG_Short, got %d", typeId)
	}

	name, err := r.readString()
	if err != nil {
		return "", err
	}

	*dst, err = r.readShort()
	return name, err
}

func (r StaticReader) Int(dst *int32) (string, error) {
	typeId, err := r.readByte()
	if err != nil {
		return "", err
	}
	if typeId != Int {
		return "", fmt.Errorf("tried reading TAG_Int, got %d", typeId)
	}

	name, err := r.readString()
	if err != nil {
		return "", err
	}

	*dst, err = r.readInt()
	return name, err
}

func (r StaticReader) Long(dst *int64) (string, error) {
	typeId, err := r.readByte()
	if err != nil {
		return "", err
	}
	if typeId != Long {
		return "", fmt.Errorf("tried reading TAG_Long, got %d", typeId)
	}

	name, err := r.readString()
	if err != nil {
		return "", err
	}

	*dst, err = r.readLong()
	return name, err
}

func (r StaticReader) String(dst *string) (string, error) {
	typeId, err := r.readByte()
	if err != nil {
		return "", err
	}
	if typeId != String {
		return "", fmt.Errorf("tried reading TAG_String, got %d", typeId)
	}

	name, err := r.readString()
	if err != nil {
		return "", err
	}

	*dst, err = r.readString()
	return name, err
}

func (r StaticReader) Compound(dst *CompoundReader) (string, error, bool) {
	typeId, err := r.readByte()
	if err != nil {
		return "", err, false
	}
	if typeId == End {
		return "", nil, true
	}
	if typeId != Compound {
		return "", fmt.Errorf("tried reading TAG_Compound, got %d", typeId), false
	}

	name, err := r.readString()
	if err != nil {
		return "", err, false
	}

	*dst = CompoundReader{r}
	return name, err, false
}

type CompoundReader struct {
	r StaticReader
}

// super convenient
func (cr CompoundReader) ReadStringMap(tgt map[string]string) error {
	for {
		typeId, err := cr.r.readByte()
		if err != nil {
			return err
		}
		if typeId == End {
			return nil
		}
		name, err := cr.r.readString()
		if err != nil {
			return err
		}

		switch typeId {
		case String:
			str, err := cr.r.readString()
			if err != nil {
				return err
			}
			tgt[name] = str
		default:
			return fmt.Errorf("unsupported tag %d for string map", typeId)
		}
	}
}

func (cr CompoundReader) ReadAll(targets ...any) error {
	for i := 0; i < len(targets)+1; i++ { //+1 to read end
		typeId, err := cr.r.readByte()
		if err != nil {
			return err
		}
		if typeId == End {
			return nil
		}
		_, err = cr.r.readString()
		if err != nil {
			return err
		}

		switch typeId {
		case Byte:
			dst, ok := targets[i].(*byte)
			if !ok {
				return fmt.Errorf("mismatched type Byte and %T", targets[i])
			}
			*dst, err = cr.r.readByte()
			if err != nil {
				return err
			}
		case Short:
			dst, ok := targets[i].(*int16)
			if !ok {
				return fmt.Errorf("mismatched type Short and %T", targets[i])
			}
			*dst, err = cr.r.readShort()
			if err != nil {
				return err
			}
		case Int:
			dst, ok := targets[i].(*int32)
			if !ok {
				return fmt.Errorf("mismatched type Int and %T", targets[i])
			}
			*dst, err = cr.r.readInt()
			if err != nil {
				return err
			}
		case Long:
			dst, ok := targets[i].(*int64)
			if !ok {
				return fmt.Errorf("mismatched type Long and %T", targets[i])
			}
			*dst, err = cr.r.readLong()
			if err != nil {
				return err
			}
		case String:
			dst, ok := targets[i].(*string)
			if !ok {
				return fmt.Errorf("mismatched type String and %T", targets[i])
			}
			*dst, err = cr.r.readString()
			if err != nil {
				return err
			}
		case Compound:
			switch dst := targets[i].(type) {
			case []any:
				CompoundReader{r: cr.r}.ReadAll(dst)
			case map[string]string:
				CompoundReader{r: cr.r}.ReadStringMap(dst)
			default:
				return fmt.Errorf("mismatched type Compound and %T", targets[i])
			}
		case List:
			dst, ok := targets[i].(func(len int32, rd ListReader))
			if !ok {
				return fmt.Errorf("mismatched type List and %T", targets[i])
			}
			typeId, err := cr.r.readByte()
			if err != nil {
				return err
			}
			len, err := cr.r.readInt()
			if err != nil {
				return err
			}
			dst(len, ListReader{typeId: typeId, r: cr.r})
		}
	}
	return nil
}

type ListReader struct {
	typeId byte
	r      StaticReader
}

func (lr ListReader) Read(tgt any) (err error) {
	switch lr.typeId {
	case Byte:
		dst, ok := tgt.(*byte)
		if !ok {
			return fmt.Errorf("mismatched type Byte and %T", tgt)
		}
		*dst, err = lr.r.readByte()
		if err != nil {
			return err
		}
	case Short:
		dst, ok := tgt.(*int16)
		if !ok {
			return fmt.Errorf("mismatched type Short and %T", tgt)
		}
		*dst, err = lr.r.readShort()
		if err != nil {
			return err
		}
	case Int:
		dst, ok := tgt.(*int32)
		if !ok {
			return fmt.Errorf("mismatched type Int and %T", tgt)
		}
		*dst, err = lr.r.readInt()
		if err != nil {
			return err
		}
	case Long:
		dst, ok := tgt.(*int64)
		if !ok {
			return fmt.Errorf("mismatched type Long and %T", tgt)
		}
		*dst, err = lr.r.readLong()
		if err != nil {
			return err
		}
	case String:
		dst, ok := tgt.(*string)
		if !ok {
			return fmt.Errorf("mismatched type String and %T", tgt)
		}
		*dst, err = lr.r.readString()
		if err != nil {
			return err
		}
	case Compound:
		dst, ok := tgt.([]any)
		if !ok {
			return fmt.Errorf("mismatched type Compound and %T", lr)
		}
		CompoundReader{r: lr.r}.ReadAll(dst...)
	case List:
		dst, ok := tgt.(func(ListReader))
		if !ok {
			return fmt.Errorf("mismatched type List and %T", tgt)
		}
		typeId, err := lr.r.readByte()
		if err != nil {
			return err
		}
		dst(ListReader{typeId: typeId, r: lr.r})
	}
	return nil
}
