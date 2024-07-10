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

type Decoder struct {
	rd io.Reader
	dontReadRootCompoundName,
	disallowUnknownFields bool
}

func NewDecoder(rd io.Reader) *Decoder {
	return &Decoder{rd: rd}
}

func (d *Decoder) ReadRootName(v bool) {
	d.dontReadRootCompoundName = !v
}

func (d *Decoder) DisallowUnknownFields(v bool) {
	d.disallowUnknownFields = v
}

func (d *Decoder) Decode(v any) (rootName string, err error) {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Pointer {
		return "", fmt.Errorf("Decode expects a pointer")
	}
	val = val.Elem()

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

	switch val.Kind() {
	case reflect.Struct:
		if err := d.decodeCompoundStruct(val); err != nil {
			return rootName, err
		}
	case reflect.Map:
		if val.IsNil() {
			val.Set(reflect.MakeMap(val.Type()))
		}
		if err := d.decodeCompoundMap(val); err != nil {
			return rootName, err
		}
	default:
		return rootName, fmt.Errorf("Decode expects a pointer of struct/map, not %s", val.Type())
	}

	return
}

func (d *Decoder) decodeCompoundStruct(_struct reflect.Value) error {
	for {
		typeId, err := d.readByte()
		if err != nil {
			return err
		}
		if typeId == End {
			return nil
		}

		name, err := d.readString()
		if err != nil {
			return err
		}

		var fieldType *reflect.StructField

		for i := 0; i < _struct.NumField(); i++ {
			typ := _struct.Type().Field(i)

			if typ.Tag.Get("nbt") == name {
				fieldType = &typ

				break
			}
		}

		if fieldType == nil {
			ft, ok := _struct.Type().FieldByName(name)
			if !ok {
				if d.disallowUnknownFields {
					return fmt.Errorf("unknown field %s", name)
				}
			}
			fieldType = &ft
		}
		field := _struct.FieldByName(fieldType.Name)

		switch typeId {
		case Byte:
			d, err := d.readByte()
			if err != nil {
				return err
			}

			if field.IsValid() {
				switch field.Kind() {
				case reflect.Uint8:
					field.SetUint(uint64(d))
				case reflect.Bool:
					field.SetBool(*(*bool)(unsafe.Pointer(&d)))
				default:
					if reflect.TypeOf(d).AssignableTo(field.Type()) {
						field.Set(reflect.ValueOf(d))
					} else {
						return fmt.Errorf("cannot assign byte to type %s for field %s", field.Type(), name)
					}
				}
			}
		case Short:
			d, err := d.readShort()
			if err != nil {
				return err
			}

			if field.IsValid() {
				switch field.Kind() {
				case reflect.Uint16:
					field.SetUint(uint64(d))
				default:
					if reflect.TypeOf(d).AssignableTo(field.Type()) {
						field.Set(reflect.ValueOf(d))
					} else {
						return fmt.Errorf("cannot assign short to type %s for field %s", field.Type(), name)
					}
				}
			}
		case Int:
			d, err := d.readInt()
			if err != nil {
				return err
			}

			if field.IsValid() {
				switch field.Kind() {
				case reflect.Uint32:
					field.SetUint(uint64(d))
				default:
					if reflect.TypeOf(d).AssignableTo(field.Type()) {
						field.Set(reflect.ValueOf(d))
					} else {
						return fmt.Errorf("cannot assign int to type %s for field %s", field.Type(), name)
					}
				}
			}
		case Long:
			d, err := d.readLong()
			if err != nil {
				return err
			}

			if field.IsValid() {
				switch field.Kind() {
				case reflect.Uint64:
					field.SetUint(uint64(d))
				default:
					if reflect.TypeOf(d).AssignableTo(field.Type()) {
						field.Set(reflect.ValueOf(d))
					} else {
						return fmt.Errorf("cannot assign long to type %s for field %s", field.Type(), name)
					}
				}
			}
		case String:
			d, err := d.readString()
			if err != nil {
				return err
			}

			if field.IsValid() {
				if reflect.TypeOf(d).AssignableTo(field.Type()) {
					field.Set(reflect.ValueOf(d))
				} else {
					return fmt.Errorf("cannot assign string to type %s for field %s", field.Type(), name)
				}
			}
		case Float:
			d, err := d.readFloat()
			if err != nil {
				return err
			}

			if field.IsValid() {
				if reflect.TypeOf(d).AssignableTo(field.Type()) {
					field.Set(reflect.ValueOf(d))
				} else {
					return fmt.Errorf("cannot assign float to type %s for field %s", field.Type(), name)
				}
			}
		case Double:
			d, err := d.readDouble()
			if err != nil {
				return err
			}

			if field.IsValid() {
				if reflect.TypeOf(d).AssignableTo(field.Type()) {
					field.Set(reflect.ValueOf(d))
				} else {
					return fmt.Errorf("cannot assign double to type %s for field %s", field.Type(), name)
				}
			}
		case ByteArray:
			data, err := d.readByteArray()
			if err != nil {
				return err
			}

			if field.IsValid() {
				switch field.Kind() {
				case reflect.Slice:
					switch field.Type().Elem().Kind() {
					case reflect.Int8:
						field.Set(reflect.ValueOf(*(*[]int8)(unsafe.Pointer(&data))))
					default:
						field.Set(reflect.ValueOf(*(*[]byte)(unsafe.Pointer(&data))))
					}
				case reflect.Array:
					switch field.Type().Elem().Kind() {
					case reflect.Int8:
						d := *(*[]int8)(unsafe.Pointer(&data))
						field.Set(reflect.ValueOf(d).Convert(field.Type()))
					default:
						field.Set(reflect.ValueOf(data).Convert(field.Type()))
					}
				default:
					if reflect.TypeOf(d).AssignableTo(field.Type()) {
						field.Set(reflect.ValueOf(data))
					} else {
						return fmt.Errorf("cannot assign byte array to type %s for field %s", field.Type(), name)
					}
				}
			}
		case IntArray:
			d, err := d.readIntArray()
			if err != nil {
				return err
			}

			if field.IsValid() {
				switch field.Kind() {
				case reflect.Slice:
					switch field.Type().Elem().Kind() {
					case reflect.Int32:
						field.Set(reflect.ValueOf(*(*[]int32)(unsafe.Pointer(&d))))
					}
				default:
					if reflect.TypeOf(d).AssignableTo(field.Type()) {
						field.Set(reflect.ValueOf(d))
					} else {
						return fmt.Errorf("cannot assign int array to type %s for field %s", field.Type(), name)
					}
				}
			}
		case LongArray:
			d, err := d.readLongArray()
			if err != nil {
				return err
			}

			if field.IsValid() {
				switch field.Kind() {
				case reflect.Slice:
					switch field.Type().Elem().Kind() {
					case reflect.Int64:
						field.Set(reflect.ValueOf(*(*[]int64)(unsafe.Pointer(&d))))
					}
				default:
					if reflect.TypeOf(d).AssignableTo(field.Type()) {
						field.Set(reflect.ValueOf(d))
					} else {
						return fmt.Errorf("cannot assign long array to type %s for field %s", field.Type(), name)
					}
				}
			}
		case List:
			if field.IsValid() {
				switch field.Kind() {
				case reflect.Slice, reflect.Array:
					if err := d.decodeList(field); err != nil {
						return err
					}
				default:
					return fmt.Errorf("cannot assign list to type %s for field %s", field.Type(), name)
				}
			} else {
				if err := d._decodeList(); err != nil {
					return err
				}
			}
		case Compound:
			if field.IsValid() {
				switch field.Kind() {
				case reflect.Struct:
					if err := d.decodeCompoundStruct(field); err != nil {
						return err
					}
				case reflect.Map:

					if field.IsNil() {
						field.Set(reflect.MakeMap(field.Type()))
					}
					if err := d.decodeCompoundMap(field); err != nil {
						return err
					}
				case reflect.Interface:
					if field.NumMethod() == 0 {
						field.Set(reflect.MakeMap(reflect.TypeOf(map[string]any{})))

						if err := d.decodeCompoundMap(field.Elem()); err != nil {
							return err
						}
						continue
					}
					fallthrough
				default:
					return fmt.Errorf("cannot assign compound to type %s for field %s", field.Type(), name)
				}
			} else {
				if err := d.decodeCompound(); err != nil {
					return err
				}
			}
		}
	}
}

func (d *Decoder) decodeCompoundMap(_map reflect.Value) error {
	for {
		typeId, err := d.readByte()
		if err != nil {
			return err
		}
		if typeId == End {
			return nil
		}

		name, err := d.readString()
		if err != nil {
			return err
		}

		nameVal := reflect.ValueOf(name)
		switch typeId {
		case Byte:
			d, err := d.readByte()
			if err != nil {
				return err
			}

			switch _map.Type().Elem().Kind() {
			case reflect.Uint8:
				_map.SetMapIndex(nameVal, reflect.ValueOf(uint8(d)))
			default:
				if reflect.TypeOf(d).AssignableTo(_map.Type().Elem()) {
					_map.SetMapIndex(nameVal, reflect.ValueOf(d))
				} else {
					return fmt.Errorf("cannot assign byte to type %s for field %s", _map.Type().Elem(), name)
				}
			}
		case Short:
			d, err := d.readShort()
			if err != nil {
				return err
			}

			switch _map.Type().Elem().Kind() {
			case reflect.Uint16:
				_map.SetMapIndex(nameVal, reflect.ValueOf(uint16(d)))
			default:
				if reflect.TypeOf(d).AssignableTo(_map.Type().Elem()) {
					_map.SetMapIndex(nameVal, reflect.ValueOf(d))
				} else {
					return fmt.Errorf("cannot assign short to type %s for field %s", _map.Type().Elem(), name)
				}
			}
		case Int:
			d, err := d.readInt()
			if err != nil {
				return err
			}

			switch _map.Type().Elem().Kind() {
			case reflect.Uint32:
				_map.SetMapIndex(nameVal, reflect.ValueOf(uint32(d)))
			default:
				if reflect.TypeOf(d).AssignableTo(_map.Type().Elem()) {
					_map.SetMapIndex(nameVal, reflect.ValueOf(d))
				} else {
					return fmt.Errorf("cannot assign int to type %s for field %s", _map.Type().Elem(), name)
				}
			}
		case Long:
			d, err := d.readLong()
			if err != nil {
				return err
			}

			switch _map.Type().Elem().Kind() {
			case reflect.Uint64:
				_map.SetMapIndex(nameVal, reflect.ValueOf(uint64(d)))
			default:
				if reflect.TypeOf(d).AssignableTo(_map.Type().Elem()) {
					_map.SetMapIndex(nameVal, reflect.ValueOf(d))
				} else {
					return fmt.Errorf("cannot assign long to type %s for field %s", _map.Type().Elem(), name)
				}
			}
		case String:
			d, err := d.readString()
			if err != nil {
				return err
			}

			if reflect.TypeOf(d).AssignableTo(_map.Type().Elem()) {
				_map.SetMapIndex(nameVal, reflect.ValueOf(d))
			} else {
				return fmt.Errorf("cannot assign long to type %s for field %s", _map.Type().Elem(), name)
			}
		case Float:
			d, err := d.readFloat()
			if err != nil {
				return err
			}

			if reflect.TypeOf(d).AssignableTo(_map.Type().Elem()) {
				_map.SetMapIndex(nameVal, reflect.ValueOf(d))
			} else {
				return fmt.Errorf("cannot assign float to type %s for field %s", _map.Type().Elem(), name)
			}
		case Double:
			d, err := d.readDouble()
			if err != nil {
				return err
			}

			if reflect.TypeOf(d).AssignableTo(_map.Type().Elem()) {
				_map.SetMapIndex(nameVal, reflect.ValueOf(d))
			} else {
				return fmt.Errorf("cannot assign float to type %s for field %s", _map.Type().Elem(), name)
			}
		case ByteArray:
			d, err := d.readByteArray()
			if err != nil {
				return err
			}

			switch _map.Type().Elem().Kind() {
			case reflect.Slice:
				switch _map.Type().Elem().Elem().Kind() {
				case reflect.Int8:
					_map.SetMapIndex(nameVal, reflect.ValueOf(*(*[]int8)(unsafe.Pointer(&d))))
				}
			default:
				if reflect.TypeOf(d).AssignableTo(_map.Type().Elem()) {
					_map.SetMapIndex(nameVal, reflect.ValueOf(d))
				} else {
					return fmt.Errorf("cannot assign byte array to type %s for field %s", _map.Type().Elem(), name)
				}
			}
		case IntArray:
			d, err := d.readIntArray()
			if err != nil {
				return err
			}

			switch _map.Type().Elem().Kind() {
			case reflect.Slice:
				switch _map.Type().Elem().Elem().Kind() {
				case reflect.Uint32:
					_map.SetMapIndex(nameVal, reflect.ValueOf(*(*[]uint32)(unsafe.Pointer(&d))))
				}
			default:
				if reflect.TypeOf(d).AssignableTo(_map.Type().Elem()) {
					_map.SetMapIndex(nameVal, reflect.ValueOf(d))
				} else {
					return fmt.Errorf("cannot assign int array to type %s for field %s", _map.Type().Elem(), name)
				}
			}
		case LongArray:
			d, err := d.readLongArray()
			if err != nil {
				return err
			}

			switch _map.Type().Elem().Kind() {
			case reflect.Slice:
				switch _map.Type().Elem().Elem().Kind() {
				case reflect.Uint64:
					_map.SetMapIndex(nameVal, reflect.ValueOf(*(*[]uint64)(unsafe.Pointer(&d))))
				}
			default:
				if reflect.TypeOf(d).AssignableTo(_map.Type().Elem()) {
					_map.SetMapIndex(nameVal, reflect.ValueOf(d))
				} else {
					return fmt.Errorf("cannot assign long array to type %s for field %s", _map.Type().Elem(), name)
				}
			}
		case List:
			switch _map.Type().Elem().Kind() {
			case reflect.Slice:
				s := reflect.MakeSlice(_map.Type().Elem(), 0, 0)
				d.decodeList(s)
				_map.SetMapIndex(nameVal, s)
			case reflect.Array:
				s := reflect.New(_map.Type().Elem()).Elem()
				d.decodeList(s)
				_map.SetMapIndex(nameVal, s)
			}
		case Compound:
			switch _map.Type().Elem().Kind() {
			case reflect.Map:
				s := reflect.MakeMap(_map.Type().Elem())
				if err := d.decodeCompoundMap(s); err != nil {
					return err
				}
				_map.SetMapIndex(nameVal, s)
			case reflect.Struct:
				s := reflect.New(_map.Type().Elem()).Elem()
				if err := d.decodeCompoundStruct(s); err != nil {
					return err
				}
				_map.SetMapIndex(nameVal, s)
			case reflect.Interface:
				if _map.Type().Elem().NumMethod() == 0 {
					s := reflect.MakeMap(reflect.TypeOf(map[string]any{}))

					if err := d.decodeCompoundMap(s); err != nil {
						return err
					}
					_map.SetMapIndex(nameVal, s)
					continue
				}
				fallthrough
			default:
				return fmt.Errorf("cannot assign compound to type %s on field %s", _map.Type().Elem(), name)
			}
		}
	}
}

func (d *Decoder) decodeList(list reflect.Value) error {
	typeId, err := d.readByte()
	if err != nil {
		return err
	}
	length, err := d.readInt()
	if err != nil {
		return err
	}
	if list.Len() < int(length) {
		switch list.Kind() {
		case reflect.Slice:
			list.Set(reflect.AppendSlice(list, reflect.MakeSlice(list.Type(), int(length)-list.Len(), int(length)-list.Len())))
		case reflect.Array:
			return fmt.Errorf("list of size %d is too big for array %s", length, list.Type())
		}
	}

	for i := 0; i < int(length); i++ {
		switch typeId {
		case Byte:
			d, err := d.readByte()
			if err != nil {
				return err
			}

			switch list.Type().Elem().Kind() {
			case reflect.Uint8:
				list.Index(i).Set(reflect.ValueOf(uint8(d)))
			default:
				if reflect.TypeOf(d).AssignableTo(list.Type().Elem()) {
					list.Index(i).Set(reflect.ValueOf(d))
				} else {
					return fmt.Errorf("cannot assign byte to type %s for index %d", list.Type().Elem(), i)
				}
			}
		case Short:
			d, err := d.readShort()
			if err != nil {
				return err
			}

			switch list.Type().Elem().Kind() {
			case reflect.Uint16:
				list.Index(i).Set(reflect.ValueOf(uint16(d)))
			default:
				if reflect.TypeOf(d).AssignableTo(list.Type().Elem()) {
					list.Index(i).Set(reflect.ValueOf(d))
				} else {
					return fmt.Errorf("cannot assign short to type %s for index %d", list.Type().Elem(), i)
				}
			}
		case Int:
			d, err := d.readInt()
			if err != nil {
				return err
			}

			switch list.Type().Elem().Kind() {
			case reflect.Uint32:
				list.Index(i).Set(reflect.ValueOf(uint32(d)))
			default:
				if reflect.TypeOf(d).AssignableTo(list.Type().Elem()) {
					list.Index(i).Set(reflect.ValueOf(d))
				} else {
					return fmt.Errorf("cannot assign int to type %s for index %d", list.Type().Elem(), i)
				}
			}
		case Long:
			d, err := d.readLong()
			if err != nil {
				return err
			}

			switch list.Type().Elem().Kind() {
			case reflect.Uint64:
				list.Index(i).Set(reflect.ValueOf(uint64(d)))
			default:
				if reflect.TypeOf(d).AssignableTo(list.Type().Elem()) {
					list.Index(i).Set(reflect.ValueOf(d))
				} else {
					return fmt.Errorf("cannot assign long to type %s for index %d", list.Type().Elem(), i)
				}
			}
		case Float:
			d, err := d.readFloat()
			if err != nil {
				return err
			}

			if reflect.TypeOf(d).AssignableTo(list.Type().Elem()) {
				list.Index(i).Set(reflect.ValueOf(d))
			} else {
				return fmt.Errorf("cannot assign float to type %s for index %d", list.Type().Elem(), i)
			}
		case Double:
			d, err := d.readDouble()
			if err != nil {
				return err
			}

			if reflect.TypeOf(d).AssignableTo(list.Type().Elem()) {
				list.Index(i).Set(reflect.ValueOf(d))
			} else {
				return fmt.Errorf("cannot assign double to type %s for index %d", list.Type().Elem(), i)
			}
		case String:
			d, err := d.readString()
			if err != nil {
				return err
			}

			if reflect.TypeOf(d).AssignableTo(list.Type().Elem()) {
				list.Index(i).Set(reflect.ValueOf(d))
			} else {
				return fmt.Errorf("cannot assign string to type %s for index %d", list.Type().Elem(), i)
			}
		case Compound:
			switch list.Type().Elem().Kind() {
			case reflect.Struct:
				if err := d.decodeCompoundStruct(list.Index(i)); err != nil {
					return err
				}
			case reflect.Map:
				list.Index(i).Set(reflect.MakeMap(list.Type().Elem()))
				if err := d.decodeCompoundMap(list.Index(i)); err != nil {
					return err
				}
			default:
				return fmt.Errorf("cannot assign compound to type %s on index %d", list.Type().Elem(), i)
			}
		case List:
			switch list.Type().Elem().Kind() {
			case reflect.Slice:
				list.Index(i).Set(reflect.MakeSlice(list.Type().Elem(), 0, 0))
				fallthrough
			case reflect.Array:
				if err := d.decodeList(list.Index(i)); err != nil {
					return err
				}
			default:
				return fmt.Errorf("cannot assign list to type %s on index %d", list.Type().Elem(), i)
			}
		}
	}

	return nil
}

func (d *Decoder) decodeCompound() error {
	for {
		typeId, err := d.readByte()
		if err != nil {
			return err
		}

		if typeId == End {
			return nil
		}
		_, err = d.readString()
		if err != nil {
			return err
		}

		switch typeId {
		case Byte:
			_, err := d.readByte()
			if err != nil {
				return err
			}
		case Short:
			_, err := d.readShort()
			if err != nil {
				return err
			}
		case Int:
			_, err := d.readInt()
			if err != nil {
				return err
			}
		case Long:
			_, err := d.readLong()
			if err != nil {
				return err
			}
		case Float:
			_, err := d.readFloat()
			if err != nil {
				return err
			}
		case Double:
			_, err := d.readDouble()
			if err != nil {
				return err
			}
		case ByteArray:
			_, err := d.readByteArray()
			if err != nil {
				return err
			}
		case IntArray:
			_, err := d.readIntArray()
			if err != nil {
				return err
			}
		case LongArray:
			_, err := d.readLongArray()
			if err != nil {
				return err
			}
		case String:
			_, err := d.readString()
			if err != nil {
				return err
			}
		case Compound:
			if err := d.decodeCompound(); err != nil {
				return err
			}
		case List:
			if err := d._decodeList(); err != nil {
				return err
			}
		}
	}
}

func (d *Decoder) _decodeList() error {
	typeId, err := d.readByte()
	if err != nil {
		return err
	}
	length, err := d.readInt()
	if err != nil {
		return err
	}
	for i := 0; i < int(length); i++ {
		switch typeId {
		case Byte:
			if _, err := d.readByte(); err != nil {
				return err
			}
		case Short:
			if _, err := d.readShort(); err != nil {
				return err
			}
		case Int:
			if _, err := d.readInt(); err != nil {
				return err
			}
		case Long:
			if _, err := d.readLong(); err != nil {
				return err
			}
		case Float:
			if _, err := d.readFloat(); err != nil {
				return err
			}
		case Double:
			if _, err := d.readDouble(); err != nil {
				return err
			}
		case String:
			if _, err := d.readString(); err != nil {
				return err
			}
		case List:
			if err := d._decodeList(); err != nil {
				return err
			}
		case Compound:
			if err := d.decodeCompound(); err != nil {
				return err
			}
		case ByteArray:
			if _, err := d.readByteArray(); err != nil {
				return err
			}
		case IntArray:
			if _, err := d.readIntArray(); err != nil {
				return err
			}
		case LongArray:
			if _, err := d.readLongArray(); err != nil {
				return err
			}
		}
	}
	return nil
}

func (d *Decoder) readByte() (int8, error) {
	var data [1]byte
	_, err := d.rd.Read(data[:])
	return int8(data[0]), err
}

func (d *Decoder) readShort() (int16, error) {
	var data [2]byte
	_, err := d.rd.Read(data[:])

	return int16(data[0])<<8 | int16(data[1]), err
}

func (d *Decoder) readInt() (int32, error) {
	var data [4]byte
	_, err := d.rd.Read(data[:])

	return int32(data[0])<<24 | int32(data[1])<<16 | int32(data[2])<<8 | int32(data[3]), err
}

func (d *Decoder) readLong() (int64, error) {
	var data [8]byte
	_, err := d.rd.Read(data[:])

	return int64(data[0])<<56 | int64(data[1])<<48 | int64(data[2])<<40 | int64(data[3])<<32 | int64(data[4])<<24 | int64(data[5])<<16 | int64(data[6])<<8 | int64(data[7]), err
}

func (d *Decoder) readFloat() (float32, error) {
	i, err := d.readInt()
	return math.Float32frombits(uint32(i)), err
}

func (d *Decoder) readDouble() (float64, error) {
	i, err := d.readLong()
	return math.Float64frombits(uint64(i)), err
}

func (d *Decoder) readString() (string, error) {
	l, err := d.readShort()

	if err != nil {
		return "", err
	}

	var data = make([]byte, l)
	_, err = d.rd.Read(data)
	return string(data), err
}

func (d *Decoder) readByteArray() ([]byte, error) {
	l, err := d.readInt()
	if err != nil {
		return nil, err
	}
	var data = make([]byte, l)
	_, err = d.rd.Read(data)

	return data, err
}

func (d *Decoder) readIntArray() ([]int32, error) {
	l, err := d.readInt()
	if err != nil {
		return nil, err
	}
	var data = make([]byte, l*4)
	_, err = d.rd.Read(data)
	if err != nil {
		return nil, err
	}

	var sl = make([]int32, l)

	for i := range sl {
		sl[i] = int32(data[i*4+0])<<24 | int32(data[i*4+1])<<16 | int32(data[i*4+2])<<8 | int32(data[i*4+3])
	}
	return sl, nil
}

func (d *Decoder) readLongArray() ([]int64, error) {
	l, err := d.readInt()
	if err != nil {
		return nil, err
	}
	var data = make([]byte, l*8)
	_, err = d.rd.Read(data)

	if err != nil {
		return nil, err
	}

	var sl = make([]int64, l)

	for i := range sl {
		sl[i] = int64(data[i*8+0])<<56 | int64(data[i*8+1])<<48 | int64(data[i*8+2])<<40 | int64(data[i*8+3])<<32 | int64(data[i*8+4])<<24 | int64(data[i*8+5])<<16 | int64(data[i*8+6])<<8 | int64(data[i*8+7])
	}
	return sl, nil
}
