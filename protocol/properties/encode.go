package properties

import (
	"fmt"
	"io"
	"reflect"
	"strings"
	"unsafe"
)

// Marshal encodes the properties file
func Marshal(dst io.Writer, p any) error {
	val := reflect.ValueOf(p)
	switch val.Kind() {
	case reflect.Struct:
		return encodePropsStruct(dst, val)
	default:
		return fmt.Errorf("Marshal excepts a struct or map, not %s", val.Kind())
	}
}

func encodePropsStruct(dst io.Writer, v reflect.Value) error {
	vt := v.Type()
	for i := 0; i < v.NumField(); i++ {
		tf := vt.Field(i)
		vf := v.Field(i)
		name, ok, omitempty := name(tf)
		if !ok {
			continue
		}
		if omitempty && vf.IsZero() {
			continue
		}

		if i != 0 {
			if err := writeString(dst, "\n"); err != nil {
				return err
			}
		}

		if err := writeString(dst, name+"="); err != nil {
			return err
		}

		switch vf.Kind() {
		case reflect.String, reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if err := writeString(dst, fmt.Sprint(vf.Interface())); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unsupported kind %s for encode struct", vf.Kind())
		}
	}

	return nil
}

func name(tf reflect.StructField) (n string, ok bool, omitempty bool) {
	if !tf.IsExported() {
		return "", false, false
	}
	name := tf.Name

	propName, ok := tf.Tag.Lookup("properties")
	if !ok {
		return name, true, false
	}
	if propName == "-" {
		return name, false, false
	}
	name = propName

	i := strings.Index(name, ",omitempty")
	if i == -1 {
		return name, true, false
	}

	name = name[:i]
	if name == "" {
		name = tf.Name
	}

	return name, true, true
}

func writeString(w io.Writer, str string) error {
	_, err := w.Write(unsafe.Slice(unsafe.StringData(str), len(str)))

	return err
}
