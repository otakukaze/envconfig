package envconfig

import "reflect"

func Export(i interface{}) []string {
	t := reflect.ValueOf(i)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	envTag := make([]string, 0)

	loop := t.NumField()
	for i := 0; i < loop; i++ {
		f := t.Field(i)
		if f.Kind() == reflect.Struct {
			envTag = append(envTag, Export(f.Addr().Interface())...)
			continue
		}

		tag := t.Type().Field(i).Tag.Get("env")
		if tag != "" {
			envTag = append(envTag, tag)
		}
	}

	return envTag
}
