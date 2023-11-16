package toml

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
)

func Unmarshal(d []byte, val any) error {
	buf := bytes.NewBuffer(d)
	e := NewDecoder(buf).Decode(val)
	return e
}

type Decoder struct {
	buf io.Reader
}

func NewDecoder(buf io.Reader) *Decoder {
	return &Decoder{buf}
}

func (d *Decoder) Decode(val any) error {
	v := reflect.ValueOf(val)
	if v.Kind() != reflect.Pointer {
		return ErrPointerRequired
	}
	b, _ := io.ReadAll(d.buf)
	t := string(b)
	if t[len(t)-1] != '\n' {
		t += "\n"
	}
	var key strings.Builder
	var value strings.Builder
	var w bool

	for _, i := range t {
		switch i {
		case '\n':
			if e := d.decode(v.Elem(), strings.TrimSpace(key.String()), strings.TrimSpace(value.String())); e != nil {
				return e
			}
			key.Reset()
			value.Reset()
			w = false
		case '=':
			w = true
		default:
			if w {
				value.WriteRune(i)
			} else {
				key.WriteRune(i)
			}
		}
	}
	return nil
}

func (d *Decoder) decode(v reflect.Value, key, value string) error {
	switch v.Kind() {
	case reflect.Struct:
		f, ok := d.findField(v, key)
		if !ok {
			return nil
		}
		switch f.Kind() {
		case reflect.String:
			{
				f.SetString(strings.TrimPrefix(strings.TrimSuffix(value, "'"), "'"))
			}
		case reflect.Bool:
			{
				b, err := strconv.ParseBool(value)
				if err != nil {
					return fmt.Errorf("failed to unmarshal field %s into type bool", key)
				}
				f.SetBool(b)
			}
		case reflect.Int:
			i, err := strconv.Atoi(value)
			if err != nil {
				return fmt.Errorf("failed to unmarshal field %s into type int", key)
			}
			f.SetInt(int64(i))
		case reflect.Int8:
			i, err := strconv.ParseInt(value, 10, 8)
			if err != nil {
				return fmt.Errorf("failed to unmarshal field %s into type int8", key)
			}
			f.SetInt(i)
		case reflect.Int16:
			i, err := strconv.ParseInt(value, 10, 16)
			if err != nil {
				return fmt.Errorf("failed to unmarshal field %s into type int16", key)
			}
			f.SetInt(i)
		case reflect.Int32:
			i, err := strconv.ParseInt(value, 10, 32)
			if err != nil {
				return fmt.Errorf("failed to unmarshal field %s into type int32", key)
			}
			f.SetInt(i)
		case reflect.Int64:
			i, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return fmt.Errorf("failed to unmarshal field %s into type int64", key)
			}
			f.SetInt(i)
		case reflect.Uint:
			i, err := strconv.ParseUint(value, 10, 0)
			if err != nil {
				return fmt.Errorf("failed to unmarshal field %s into type uint", key)
			}
			f.SetUint(i)
		case reflect.Uint8:
			i, err := strconv.ParseUint(value, 10, 8)
			if err != nil {
				return fmt.Errorf("failed to unmarshal field %s into type uint8", key)
			}
			f.SetUint(i)
		case reflect.Uint16:
			i, err := strconv.ParseUint(value, 10, 16)
			if err != nil {
				return fmt.Errorf("failed to unmarshal field %s into type uint16", key)
			}
			f.SetUint(i)
		case reflect.Uint32:
			i, err := strconv.ParseUint(value, 10, 32)
			if err != nil {
				return fmt.Errorf("failed to unmarshal field %s into type uint32", key)
			}
			f.SetUint(i)
		case reflect.Uint64:
			i, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return fmt.Errorf("failed to unmarshal field %s into type uint64", key)
			}
			f.SetUint(i)
		}
	default:
		return ErrUnknownType
	}
	return nil
}

func (d *Decoder) findField(s reflect.Value, key string) (reflect.Value, bool) {
	t := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		tf := t.Field(i)
		if tf.Tag.Get("toml") == key {
			return f, true
		}
	}

	f := s.FieldByName(key)
	_, ok := s.Type().FieldByName(key)
	return f, ok
}
