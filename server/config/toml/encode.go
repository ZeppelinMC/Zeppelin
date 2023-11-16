package toml

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"reflect"
)

var ErrUnknownType = errors.New("unknown object type")
var ErrPointerRequired = errors.New("decoder expects a pointer value")
var ErrInvalidSyntax = errors.New("invalid syntax")

func Marshal(val any) ([]byte, error) {
	var buf bytes.Buffer
	e := NewEncoder(&buf).Encode(val)
	return buf.Bytes(), e
}

type Encoder struct {
	buf io.Writer
}

func NewEncoder(buf io.Writer) *Encoder {
	return &Encoder{buf}
}

func (e *Encoder) Encode(val any) error {
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.Struct:
		t := v.Type()
		for i := 0; i < v.NumField(); i++ {
			f := v.Field(i)
			tf := t.Field(i)
			if !tf.IsExported() {
				continue
			}

			name := tf.Tag.Get("toml")
			if name == "" {
				name = tf.Name
			}
			e.encode(name, f)
		}
	case reflect.Map:
		for _, k := range v.MapKeys() {
			f := v.MapIndex(k)
			e.encode(k.String(), f)
		}
	default:
		return ErrUnknownType
	}
	return nil
}

func (e *Encoder) write(val reflect.Value) {
	switch val.Kind() {
	case reflect.String:
		fmt.Fprintf(e.buf, "'%s'", val.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		fmt.Fprint(e.buf, val.Int())
	case reflect.Bool:
		fmt.Fprint(e.buf, val.Bool())
	}
}

func (e *Encoder) encode(name string, val reflect.Value) {
	switch val.Kind() {
	case reflect.String:
		fmt.Fprintf(e.buf, "%s = '%s'", name, val.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		fmt.Fprintf(e.buf, "%s = %d", name, val.Int())
	case reflect.Bool:
		fmt.Fprintf(e.buf, "%s = %v", name, val.Bool())
	case reflect.Slice:
		fmt.Fprintf(e.buf, "%s = [", name)
		for i := 0; i < val.Len(); i++ {
			e.write(val.Index(i))
		}
		fmt.Fprint(e.buf, "]")
	case reflect.Struct:
		fmt.Fprintf(e.buf, "[%s]\n", name)
		t := val.Type()
		for i := 0; i < val.NumField(); i++ {
			f := val.Field(i)
			tf := t.Field(i)
			if !tf.IsExported() {
				continue
			}

			name := tf.Tag.Get("toml")
			if name == "" {
				name = tf.Name
			}
			e.encode(name, f)
		}
	case reflect.Map:
		fmt.Fprintf(e.buf, "\n[%s]\n", name)
		for _, k := range val.MapKeys() {
			f := val.MapIndex(k)
			e.encode(k.String(), f)
		}
	}
	fmt.Fprintln(e.buf)
}
