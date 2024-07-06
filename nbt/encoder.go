package nbt

import (
	"encoding/binary"
	"fmt"
	"io"
	"reflect"
	"unsafe"
)

type Encoder struct {
	w io.Writer
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w}
}

func (e *Encoder) Encode(name string, v any) error {
	if err := e.writeByte(Compound); err != nil {
		return err
	}
	if err := e.writeString(name); err != nil {
		return err
	}
	return e.encodeCompound(v)
}

func (e *Encoder) writeBytes(b ...byte) error {
	_, err := e.w.Write(b)
	return err
}

func (e *Encoder) writeByte(b int8) error {
	return e.writeBytes(byte(b))
}

func (e *Encoder) writeShort(s int16) error {
	return binary.Write(e.w, binary.BigEndian, s)
}

func (e *Encoder) writeInt(i int32) error {
	return binary.Write(e.w, binary.BigEndian, i)
}

func (e *Encoder) writeLong(l int64) error {
	return binary.Write(e.w, binary.BigEndian, l)
}

func (e *Encoder) writeFloat(f float32) error {
	return binary.Write(e.w, binary.BigEndian, f)
}

func (e *Encoder) writeDouble(d float64) error {
	return binary.Write(e.w, binary.BigEndian, d)
}

func (e *Encoder) writeByteArray(ba []int8) error {
	if err := e.writeInt(int32(len(ba))); err != nil {
		return err
	}
	return e.writeBytes(*(*[]byte)(unsafe.Pointer(&ba))...)
}

func (e *Encoder) writeIntArray(il []int32) error {
	if err := e.writeInt(int32(len(il))); err != nil {
		return err
	}
	return binary.Write(e.w, binary.BigEndian, il)
}

func (e *Encoder) writeLongArray(ll []int64) error {
	if err := e.writeInt(int32(len(ll))); err != nil {
		return err
	}
	return binary.Write(e.w, binary.BigEndian, ll)
}

func (e *Encoder) writeString(s string) error {
	if err := e.writeShort(int16(len(s))); err != nil {
		return err
	}
	return e.writeBytes([]byte(s)...)
}

func (e *Encoder) encodeCompound(v any) error {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	switch val.Kind() {
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			fieldType := val.Type().Field(i)
			name := fieldType.Name
			tag := fieldType.Tag.Get("nbt")
			if tag != "" {
				name = tag
			}
			if err := e.encodeField(name, field); err != nil {
				return err
			}
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			name := key.String()
			if err := e.encodeField(name, val.MapIndex(key)); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported compound type: %s", val.Type())
	}
	return e.writeByte(End)
}

func (e *Encoder) encodeField(name string, val reflect.Value) error {
	switch val.Kind() {
	case reflect.Int8:
		if err := e.writeByte(Byte); err != nil {
			return err
		}
		if err := e.writeString(name); err != nil {
			return err
		}
		return e.writeByte(int8(val.Int()))
	case reflect.Int16:
		if err := e.writeByte(Short); err != nil {
			return err
		}
		if err := e.writeString(name); err != nil {
			return err
		}
		return e.writeShort(int16(val.Int()))
	case reflect.Int32:
		if err := e.writeByte(Int); err != nil {
			return err
		}
		if err := e.writeString(name); err != nil {
			return err
		}
		return e.writeInt(int32(val.Int()))
	case reflect.Int64:
		if err := e.writeByte(Long); err != nil {
			return err
		}
		if err := e.writeString(name); err != nil {
			return err
		}
		return e.writeLong(val.Int())
	case reflect.Float32:
		if err := e.writeByte(Float); err != nil {
			return err
		}
		if err := e.writeString(name); err != nil {
			return err
		}
		return e.writeFloat(float32(val.Float()))
	case reflect.Float64:
		if err := e.writeByte(Double); err != nil {
			return err
		}
		if err := e.writeString(name); err != nil {
			return err
		}
		return e.writeDouble(val.Float())
	case reflect.String:
		if err := e.writeByte(String); err != nil {
			return err
		}
		if err := e.writeString(name); err != nil {
			return err
		}
		return e.writeString(val.String())
	case reflect.Slice:
		switch val.Type().Elem().Kind() {
		case reflect.Int8:
			if err := e.writeByte(ByteArray); err != nil {
				return err
			}
			if err := e.writeString(name); err != nil {
				return err
			}
			return e.writeByteArray(val.Interface().([]int8))
		case reflect.Int32:
			if err := e.writeByte(IntArray); err != nil {
				return err
			}
			if err := e.writeString(name); err != nil {
				return err
			}
			return e.writeIntArray(val.Interface().([]int32))
		case reflect.Int64:
			if err := e.writeByte(LongArray); err != nil {
				return err
			}
			if err := e.writeString(name); err != nil {
				return err
			}
			return e.writeLongArray(val.Interface().([]int64))
		default:
			if err := e.writeByte(List); err != nil {
				return err
			}
			if err := e.writeString(name); err != nil {
				return err
			}
			if val.Len() == 0 {
				if err := e.writeByte(End); err != nil {
					return err
				}
				return e.writeInt(0)
			}
			elemType := val.Index(0).Kind()
			if err := e.encodeList(val, elemType); err != nil {
				return err
			}
		}
	case reflect.Map:
		if err := e.writeByte(Compound); err != nil {
			return err
		}
		if err := e.writeString(name); err != nil {
			return err
		}
		return e.encodeCompound(val.Interface())
	case reflect.Struct:
		if err := e.writeByte(Compound); err != nil {
			return err
		}
		if err := e.writeString(name); err != nil {
			return err
		}
		return e.encodeCompound(val.Interface())
	default:
		return fmt.Errorf("unsupported field type: %s", val.Type())
	}
	return nil
}

func (e *Encoder) encodeList(val reflect.Value, elemType reflect.Kind) error {
	var elemTypeId int8
	switch elemType {
	case reflect.Int8:
		elemTypeId = Byte
	case reflect.Int16:
		elemTypeId = Short
	case reflect.Int32:
		elemTypeId = Int
	case reflect.Int64:
		elemTypeId = Long
	case reflect.Float32:
		elemTypeId = Float
	case reflect.Float64:
		elemTypeId = Double
	case reflect.String:
		elemTypeId = String
	case reflect.Slice:
		switch val.Index(0).Type().Elem().Kind() {
		case reflect.Int8:
			elemTypeId = ByteArray
		case reflect.Int32:
			elemTypeId = IntArray
		case reflect.Int64:
			elemTypeId = LongArray
		default:
			elemTypeId = List
		}
	case reflect.Map, reflect.Struct:
		elemTypeId = Compound
	default:
		return fmt.Errorf("unsupported list element type: %s", elemType)
	}

	if err := e.writeByte(elemTypeId); err != nil {
		return err
	}
	if err := e.writeInt(int32(val.Len())); err != nil {
		return err
	}

	for i := 0; i < val.Len(); i++ {
		switch elemTypeId {
		case Byte:
			if err := e.writeByte(int8(val.Index(i).Int())); err != nil {
				return err
			}
		case Short:
			if err := e.writeShort(int16(val.Index(i).Int())); err != nil {
				return err
			}
		case Int:
			if err := e.writeInt(int32(val.Index(i).Int())); err != nil {
				return err
			}
		case Long:
			if err := e.writeLong(val.Index(i).Int()); err != nil {
				return err
			}
		case Float:
			if err := e.writeFloat(float32(val.Index(i).Float())); err != nil {
				return err
			}
		case Double:
			if err := e.writeDouble(val.Index(i).Float()); err != nil {
				return err
			}
		case String:
			if err := e.writeString(val.Index(i).String()); err != nil {
				return err
			}
		case ByteArray:
			if err := e.writeByteArray(val.Index(i).Interface().([]int8)); err != nil {
				return err
			}
		case IntArray:
			if err := e.writeIntArray(val.Index(i).Interface().([]int32)); err != nil {
				return err
			}
		case LongArray:
			if err := e.writeLongArray(val.Index(i).Interface().([]int64)); err != nil {
				return err
			}
		case Compound:
			if err := e.encodeCompound(val.Index(i).Interface()); err != nil {
				return err
			}
		case List:
			if err := e.encodeList(val.Index(i), val.Index(i).Index(0).Kind()); err != nil {
				return err
			}
		}
	}
	return nil
}
