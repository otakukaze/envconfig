package envconfig

import (
	"log"
	"reflect"
)

// Parse -
func Parse(i interface{}) {
	t := reflect.ValueOf(i)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	loop := t.NumField()
	for i := 0; i < loop; i++ {
		f := t.Field(i)
		if f.Kind() == reflect.Struct {
			Parse(f.Addr().Interface())
			continue
		}

		tag := t.Type().Field(i).Tag.Get("env")
		if len(tag) == 0 {
			continue
		}

		switch f.Kind() {
		case reflect.Int:
			if data, ok := parseInt(tag); ok {
				f.SetInt(int64(data))
			}
			break
		case reflect.Int16:
			if data, ok := parseInt16(tag); ok {
				f.SetInt(int64(data))
			}
			break
		case reflect.Int32:
			if data, ok := parseInt32(tag); ok {
				f.SetInt(int64(data))
			}
			break
		case reflect.Int64:
			if data, ok := parseInt64(tag); ok {
				f.SetInt(data)
			}
			break
		case reflect.String:
			if data, ok := parseString(tag); ok {
				f.SetString(data)
			}
			break
		case reflect.Bool:
			if data, ok := parseBool(tag); ok {
				f.SetBool(data)
			}
			break
		case reflect.Float32:
			if data, ok := parseFloat32(tag); ok {
				f.Set(reflect.ValueOf(data))
			}
			break
		case reflect.Float64:
			if data, ok := parseFloat64(tag); ok {
				f.Set(reflect.ValueOf(data))
			}
			break
		case reflect.Slice:
			ftype := f.Type().Elem().Kind()
			switch ftype {
			case reflect.String:
				if val, ok := parseSliceString(tag); ok {
					f.Set(reflect.ValueOf(val))
				}
				break
			case reflect.Int:
				if val, ok := parseSliceInt(tag); ok {
					f.Set(reflect.ValueOf(val))
				}
				break
			case reflect.Int16:
				if val, ok := parseSliceInt16(tag); ok {
					f.Set(reflect.ValueOf(val))
				}
				break
			case reflect.Int32:
				if val, ok := parseSliceInt32(tag); ok {
					f.Set(reflect.ValueOf(val))
				}
				break
			case reflect.Int64:
				if val, ok := parseSliceInt64(tag); ok {
					f.Set(reflect.ValueOf(val))
				}
				break
			case reflect.Bool:
				if val, ok := parseSliceBool(tag); ok {
					f.Set(reflect.ValueOf(val))
				}
				break
			case reflect.Float32:
				if val, ok := parseSliceFloat32(tag); ok {
					f.Set(reflect.ValueOf(val))
				}
				break
			case reflect.Float64:
				if val, ok := parseSliceFloat64(tag); ok {
					f.Set(reflect.ValueOf(val))
				}
				break
			}
			break
		default:
			log.Printf("%v is not support \n", f.Kind())
			continue
		}
	}
}
