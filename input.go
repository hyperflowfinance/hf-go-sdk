package hyperflow

import (
	"encoding/json"
	"reflect"
	"strings"
)

type omitted struct{}

func marshalInput(value any) ([]byte, error) {
	return json.Marshal(normalize(reflect.ValueOf(value)))
}

func normalize(value reflect.Value) any {
	if !value.IsValid() {
		return nil
	}
	if value.CanInterface() {
		if optional, ok := value.Interface().(optionalValue); ok {
			item, set := optional.graphOptional()
			if !set {
				return omitted{}
			}
			return normalize(reflect.ValueOf(item))
		}
		if _, ok := value.Interface().(json.RawMessage); ok {
			return value.Interface()
		}
	}
	switch value.Kind() {
	case reflect.Interface, reflect.Pointer:
		if value.IsNil() {
			return nil
		}
		return normalize(value.Elem())
	case reflect.Struct:
		result := map[string]any{}
		typ := value.Type()
		for i := 0; i < value.NumField(); i++ {
			field := typ.Field(i)
			if field.PkgPath != "" {
				continue
			}
			name := strings.Split(field.Tag.Get("json"), ",")[0]
			if name == "-" {
				continue
			}
			if name == "" {
				name = field.Name
			}
			item := normalize(value.Field(i))
			if _, skip := item.(omitted); !skip {
				if field.Anonymous {
					if embedded, ok := item.(map[string]any); ok {
						for key, nested := range embedded {
							result[key] = nested
						}
						continue
					}
				}
				result[name] = item
			}
		}
		return result
	case reflect.Slice, reflect.Array:
		if value.IsNil() {
			return nil
		}
		result := make([]any, value.Len())
		for i := range result {
			item := normalize(value.Index(i))
			if _, skip := item.(omitted); skip {
				item = nil
			}
			result[i] = item
		}
		return result
	case reflect.Map:
		result := map[string]any{}
		iter := value.MapRange()
		for iter.Next() {
			result[iter.Key().String()] = normalize(iter.Value())
		}
		return result
	default:
		return value.Interface()
	}
}
