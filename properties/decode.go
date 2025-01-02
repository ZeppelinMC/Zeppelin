// Package properties provides parsing of .properties files
package properties

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Unmarshal decodes the properties file
func Unmarshal(src string, dst any) error {
	properties := strings.Split(src, "\n")
	val := reflect.ValueOf(dst)

	switch val.Kind() {
	case reflect.Pointer:
		val = val.Elem()
		switch val.Kind() {
		case reflect.Struct:
			return unmarshalStruct(properties, structMap(val))
		default:
			return fmt.Errorf("Unmarshal excepts a pointer of a struct or map, not %s", val.Kind())
		}
	default:
		return fmt.Errorf("Unmarshal excepts a pointer of a struct or map, not %s", val.Kind())
	}
}

func unmarshalStruct(props []string, v map[string]reflect.Value) error {
	for _, line := range props {
		if len(line) == 0 {
			continue //comment
		}
		if line[0] == '#' {
			continue //comment
		}
		if line[0] == '!' {
			continue //comment
		}

		var key, value string
		for i, char := range line {
			if char == '=' || char == ':' || char == ' ' {
				key = line[:i]
				if i != len(line)-2 {
					value = line[i+1:]
				}
				break
			}
		}

		if value == "" {
			continue
		}

		value = strings.TrimSpace(value)
		value = strings.ReplaceAll(value, "\\n", "\n")
		value = strings.ReplaceAll(value, "\\r", "\r")
		value = strings.ReplaceAll(value, "\\t", "\t")

		field, ok := v[key]
		if !ok {
			continue
		}

		switch field.Kind() {
		case reflect.String:
			field.SetString(value)
		case reflect.Bool:
			switch value {
			case "true":
				field.SetBool(true)
			case "false":
				field.SetBool(false)
			default:
				return fmt.Errorf("unsupported value %s for type boolean", value)
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			i, err := strconv.ParseInt(value, 10, field.Type().Bits())
			if err != nil {
				return fmt.Errorf("unsupported value %s for type integer: %v", value, err)
			}
			field.SetInt(i)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			i, err := strconv.ParseUint(value, 10, field.Type().Bits())
			if err != nil {
				return fmt.Errorf("unsupported value %s for type integer: %v", value, err)
			}
			field.SetUint(i)
		}
	}

	return nil
}

func structMap(val reflect.Value) map[string]reflect.Value {
	var sm = make(map[string]reflect.Value, val.NumField())

	vt := val.Type()
	for i := 0; i < val.NumField(); i++ {
		ft := vt.Field(i)
		fv := val.Field(i)

		if !ft.IsExported() {
			continue
		}
		name := ft.Name
		propName, ok := ft.Tag.Lookup("properties")
		if ok {
			if propName == "-" {
				continue
			}
			name = propName
			if i := strings.Index(name, ",omitempty"); i != -1 {
				name = name[:i]
			}
		}
		if name == "" {
			name = ft.Name
		}

		sm[name] = fv
	}

	return sm
}
