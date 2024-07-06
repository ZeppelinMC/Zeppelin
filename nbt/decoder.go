package nbt

import (
	"fmt"
	"io"
	"math"
	"reflect"
	"unsafe"
)

const (
	End = iota
	Byte
	Short
	Int
	Long
	Float
	Double
	ByteArray
	String
	List
	Compound
	IntArray
	LongArray
)

// Decoder doesnt work with nested maps yet
type Decoder struct {
	rd io.Reader
	dontReadRootCompoundName,
	disallowUnknownFields bool

	i int
}

func NewDecoder(rd io.Reader) *Decoder {
	return &Decoder{rd: rd}
}

func (d *Decoder) Decode(v any) (rootName string, err error) {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Pointer && val.Kind() != reflect.Map {
		return "", fmt.Errorf("Decode expects a pointer")
	}
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	typeId, err := d.readByte()
	if err != nil {
		return "", err
	}
	if typeId != Compound {
		return "", fmt.Errorf("expected a compound first element")
	}

	if !d.dontReadRootCompoundName {
		rootName, err = d.readString()
		if err != nil {
			return
		}
	}
	err = d.decodeCompound(val)

	return
}

func (d *Decoder) ReadRootName(val bool) {
	d.dontReadRootCompoundName = !val
}

func (d *Decoder) DisallowUnknownFields(val bool) {
	d.disallowUnknownFields = val
}

func (d *Decoder) readBytes(l int) ([]byte, error) {
	data := make([]byte, l)

	_, err := d.rd.Read(data)
	return data, err
}

func (d *Decoder) readByte() (int8, error) {
	data, err := d.readBytes(1)

	return int8(data[0]), err
}

func (d *Decoder) readShort() (int16, error) {
	data, err := d.readBytes(2)
	return int16(data[0])<<8 | int16(data[1]), err
}

func (d *Decoder) readInt() (int32, error) {
	data, err := d.readBytes(4)
	return int32(data[0])<<24 | int32(data[1])<<16 | int32(data[2])<<8 | int32(data[3]), err
}

func (d *Decoder) readLong() (int64, error) {
	data, err := d.readBytes(8)
	return int64(data[0])<<56 | int64(data[1])<<48 | int64(data[2])<<40 | int64(data[3])<<32 | int64(data[4])<<24 | int64(data[5])<<16 | int64(data[6])<<8 | int64(data[7]), err
}

func (d *Decoder) readFloat() (float32, error) {
	i32, err := d.readInt()

	return math.Float32frombits(uint32(i32)), err
}

func (d *Decoder) readDouble() (float64, error) {
	i64, err := d.readLong()

	return math.Float64frombits(uint64(i64)), err
}

func (d *Decoder) readByteArray() ([]int8, error) {
	length, err := d.readInt()
	if err != nil {
		return nil, err
	}

	data, err := d.readBytes(int(length))

	return *(*[]int8)(unsafe.Pointer(&data)), err
}

func (d *Decoder) readIntArray() ([]int32, error) {
	length, err := d.readInt()
	if err != nil {
		return nil, err
	}

	data, err := d.readBytes(int(length) * 4)
	sl := make([]int32, length)

	for i := 0; i < len(sl); i++ {
		sl[i] = int32(data[i*4])<<24 | int32(data[i*4+1])<<16 | int32(data[i*4+2])<<8 | int32(data[i*4+3])
	}

	return sl, err
}

func (d *Decoder) readLongArray() ([]int64, error) {
	length, err := d.readInt()
	if err != nil {
		return nil, err
	}

	data, err := d.readBytes(int(length) * 8)
	sl := make([]int64, length)

	for i := 0; i < len(sl); i++ {
		sl[i] = int64(data[i*8])<<56 | int64(data[i*8+1])<<48 | int64(data[i*8+2])<<40 | int64(data[i*8+3])<<32 | int64(data[i*8+4])<<24 | int64(data[i*8+5])<<16 | int64(data[i*8+6])<<8 | int64(data[i*8+7])
	}

	return sl, err
}

func (d *Decoder) readString() (string, error) {
	length, err := d.readShort()
	if err != nil {
		return "", err
	}

	data, err := d.readBytes(int(length))
	if err != nil {
		return "", err
	}
	var stringData = make([]rune, length)
	var finalLength int

	for i := 0; i < int(length); i++ {
		i += decodeChar(&stringData[i], data[i:])

		finalLength++
	}
	stringData = stringData[:finalLength]

	return string(stringData), nil
}

func (d *Decoder) decodeCompound(val reflect.Value) error {
	for {
		elemType, err := d.readByte()
		if err != nil {
			return err
		}
		if elemType == End {
			return nil
		}
		name, err := d.readString()
		if err != nil {
			return err
		}
		var value any
		switch elemType {
		case Byte:
			value, err = d.readByte()
			if err != nil {
				return err
			}
		case Short:
			value, err = d.readShort()
			if err != nil {
				return err
			}
		case Int:
			value, err = d.readInt()
			if err != nil {
				return err
			}
		case Long:
			value, err = d.readLong()
			if err != nil {
				return err
			}
		case Float:
			value, err = d.readFloat()
			if err != nil {
				return err
			}
		case Double:
			value, err = d.readDouble()
			if err != nil {
				return err
			}
		case String:
			value, err = d.readString()
			if err != nil {
				return err
			}
		case ByteArray:
			value, err = d.readByteArray()
			if err != nil {
				return err
			}
		case IntArray:
			value, err = d.readIntArray()
			if err != nil {
				return err
			}
		case LongArray:
			value, err = d.readLongArray()
			if err != nil {
				return err
			}
		case Compound:
			c, err := d.compoundGetCompound(val, name)
			if err != nil {
				return err
			}
			if err := d.decodeCompound(c); err != nil {
				return err
			}
			continue
		case List:
			c, err := d.compoundGetList(val, name)
			if err != nil {
				return err
			}
			if err := d.decodeList(c); err != nil {
				return err
			}
			continue
		}

		if err := d.compoundSet(val, name, value); err != nil {
			return err
		}
	}
}

func (d *Decoder) decodeList(val reflect.Value) error {
	typeId, err := d.readByte()
	if err != nil {
		return err
	}
	length, err := d.readInt()
	if err != nil {
		return err
	}
	if typeId == End && length > 0 {
		return fmt.Errorf("unexpected list of Tag_End")
	}

	// Initialize the slice if necessary
	if val.Kind() == reflect.Interface {
		val.Set(reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf((*interface{})(nil)).Elem()), int(length), int(length)))
		val = val.Elem()
	} else if val.Kind() == reflect.Slice {
		val.Set(reflect.MakeSlice(val.Type(), int(length), int(length)))
	} else if val.Kind() == reflect.Array {
		if length > int32(val.Len()) {
			return fmt.Errorf("len %d is bigger than array len %d", length, val.Len())
		}
	}

	for i := int32(0); i < length; i++ {
		var value any
		switch typeId {
		case Byte:
			value, err = d.readByte()
			if err != nil {
				return err
			}
		case Short:
			value, err = d.readShort()
			if err != nil {
				return err
			}
		case Int:
			value, err = d.readInt()
			if err != nil {
				return err
			}
		case Long:
			value, err = d.readLong()
			if err != nil {
				return err
			}
		case Float:
			value, err = d.readFloat()
			if err != nil {
				return err
			}
		case Double:
			value, err = d.readDouble()
			if err != nil {
				return err
			}
		case String:
			value, err = d.readString()
			if err != nil {
				return err
			}
		case ByteArray:
			value, err = d.readByteArray()
			if err != nil {
				return err
			}
		case IntArray:
			value, err = d.readIntArray()
			if err != nil {
				return err
			}
		case LongArray:
			value, err = d.readLongArray()
			if err != nil {
				return err
			}
		case Compound:
			c, err := d.listGetCompound(val, int(i))
			if err != nil {
				return err
			}
			if err := d.decodeCompound(c); err != nil {
				return err
			}
			continue
		case List:
			c, err := d.listGetList(val, int(i))
			if err != nil {
				return err
			}
			if err := d.decodeList(c); err != nil {
				return err
			}
			continue
		}

		if err := d.listSet(val, int(i), value); err != nil {
			return err
		}
	}
	return nil
}

func decodeChar(tgt *rune, src []byte) (itrinc int) {
	switch {
	case src[0]&0x80 == 0: //'\u0001' to '\u007F'
		*tgt = rune(src[0])
	case src[0]&0xE0 == 0xC0 && src[1]&0xC0 == 0x80: //'\u0000' and characters in the range '\u0080' to '\u07FF'
		b1 := src[0] & ((1 << 5) - 1)
		b2 := src[1] & ((1 << 6) - 1)
		*tgt = rune(b2) | rune(b1)<<5
		itrinc++
	case src[0]&0xF0 == 0xE0 && src[1]&0xC0 == 0x80 && src[2]&0xC0 == 0x80: //'\u0800' to '\uFFFF'
		b1 := src[0] & ((1 << 4) - 1)
		b2 := src[1] & ((1 << 6) - 1)
		b3 := src[2] & ((1 << 6) - 1)

		*tgt = rune(b3) | rune(b2)<<5 | rune(b1)<<10
		itrinc += 2
	}

	return
}

func (d *Decoder) compoundGetList(val reflect.Value, name string) (reflect.Value, error) {
	z := reflect.MakeSlice(reflect.TypeOf([]any{}), 0, 0)
	switch val.Kind() {
	case reflect.Struct:
		field := val.FieldByName(name)
		if !field.IsValid() {
			for i := 0; i < val.NumField(); i++ {
				if val.Type().Field(i).Tag.Get("nbt") == name {
					field = val.Field(i)
					goto cont
				}
			}
			if d.disallowUnknownFields {
				return z, fmt.Errorf("unknown field %s for struct %s", name, val)
			}
			return z, nil
		}

	cont:
		if field.Kind() != reflect.Slice && field.Kind() != reflect.Array {
			if field.Kind() == reflect.Interface {
				if field.NumMethod() != 0 {
					return z, fmt.Errorf("cannot assign list to %s for field %s", field.Type(), name)
				}
			} else {
				return z, fmt.Errorf("cannot assign list to %s for field %s", field.Type(), name)
			}
		}
		return field, nil
	case reflect.Map:
		if val.Type().Elem().Kind() != reflect.Slice && val.Type().Elem().Kind() != reflect.Array {
			if val.Type().Elem().Kind() == reflect.Interface {
				if val.Type().Elem().NumMethod() != 0 {
					return z, fmt.Errorf("cannot assign compound to %s for field %s", val.Type().Elem(), name)
				}
			} else {
				return z, fmt.Errorf("cannot assign compound to %s for field %s", val.Type().Elem(), name)
			}
		}
		v := val.MapIndex(reflect.ValueOf(name))
		if !v.IsValid() {
			val.SetMapIndex(reflect.ValueOf(name), z)
			v = val.MapIndex(reflect.ValueOf(name))
		}
		return v, nil
	}
	return z, nil
}

func (d *Decoder) compoundGetCompound(val reflect.Value, name string) (reflect.Value, error) {
	z := reflect.MakeMap(reflect.TypeOf(map[string]any{}))
	switch val.Kind() {
	case reflect.Struct:
		field := val.FieldByName(name)
		if !field.IsValid() {
			for i := 0; i < val.NumField(); i++ {
				if val.Type().Field(i).Tag.Get("nbt") == name {
					field = val.Field(i)
					goto cont
				}
			}
			if d.disallowUnknownFields {
				return z, fmt.Errorf("unknown field %s for struct %s", name, val)
			}
			return z, nil
		}
	cont:
		if field.Kind() != reflect.Map && field.Kind() != reflect.Struct {
			if field.Kind() == reflect.Interface {
				if field.NumMethod() != 0 {
					return z, fmt.Errorf("cannot assign compound to %s for field %s", field.Type(), name)
				}
			} else {
				return z, fmt.Errorf("cannot assign compound to %s for field %s", field.Type(), name)
			}
		}
		if field.Kind() == reflect.Map && field.IsNil() {
			field.Set(reflect.MakeMap(field.Type()))
		}
		return field, nil
	case reflect.Map:
		if val.Type().Elem().Kind() != reflect.Map && val.Type().Elem().Kind() != reflect.Struct {
			if val.Type().Elem().Kind() == reflect.Interface {
				if val.Type().Elem().NumMethod() != 0 {
					return z, fmt.Errorf("cannot assign compound to %s for field %s", val.Type().Elem(), name)
				}
			} else {
				return z, fmt.Errorf("cannot assign compound to %s for field %s", val.Type().Elem(), name)
			}
		}
		v := val.MapIndex(reflect.ValueOf(name))
		if !v.IsValid() {
			val.SetMapIndex(reflect.ValueOf(name), z)
			v = val.MapIndex(reflect.ValueOf(name))
		}
		return v, nil
	}
	return z, nil
}

func (d *Decoder) compoundSet(val reflect.Value, name string, value any) error {
	valueType := reflect.TypeOf(value)
	switch val.Kind() {
	case reflect.Struct:
		field := val.FieldByName(name)
		if !field.IsValid() {
			for i := 0; i < val.NumField(); i++ {
				if val.Type().Field(i).Tag.Get("nbt") == name {
					field = val.Field(i)
					goto cont
				}
			}
			if d.disallowUnknownFields {
				return fmt.Errorf("unknown field %s for struct %s", name, val)
			}
			return nil
		}
	cont:
		if !valueType.AssignableTo(field.Type()) {
			return fmt.Errorf("cannot assign %s to %s for field %s", valueType, field.Type(), name)
		}
		field.Set(reflect.ValueOf(value))
	case reflect.Map:
		if !valueType.AssignableTo(val.Type().Elem()) {
			return fmt.Errorf("cannot assign %s to %s for field %s", valueType, val.Type().Elem(), name)
		}
		if val.IsNil() {
			val.Set(reflect.MakeMap(val.Type()))
		}
		val.SetMapIndex(reflect.ValueOf(name), reflect.ValueOf(value))
	}
	return nil
}

func (d *Decoder) listGetList(val reflect.Value, index int) (reflect.Value, error) {
	z := reflect.MakeSlice(reflect.TypeOf([]any{}), 0, 0)
	field := val.Index(index)
	if field.Kind() != reflect.Slice && field.Kind() != reflect.Array {
		if field.Kind() == reflect.Interface {
			if field.NumMethod() != 0 {
				return z, fmt.Errorf("cannot assign compound to %s for element %d", field.Type(), index)
			}
		} else {
			return z, fmt.Errorf("cannot assign compound to %s for element %d", field.Type(), index)
		}
	}
	return field, nil
}

func (d *Decoder) listGetCompound(val reflect.Value, index int) (reflect.Value, error) {
	z := reflect.MakeMap(reflect.TypeOf(map[string]any{}))
	field := val.Index(index)
	if field.Kind() != reflect.Map && field.Kind() != reflect.Struct {
		if field.Kind() == reflect.Interface {
			if field.NumMethod() != 0 {
				return z, fmt.Errorf("cannot assign compound to %s for element %d", field.Type(), index)
			}
		} else {
			return z, fmt.Errorf("cannot assign compound to %s for element %d", field.Type(), index)
		}
	}
	return field, nil
}

func (d *Decoder) listSet(val reflect.Value, index int, value any) error {
	valueType := reflect.TypeOf(value)
	if !valueType.AssignableTo(val.Type().Elem()) {
		return fmt.Errorf("cannot assign %s to %s for element %d", valueType, val.Type().Elem(), index)
	}

	field := val.Index(index)
	field.Set(reflect.ValueOf(value))
	return nil
}
