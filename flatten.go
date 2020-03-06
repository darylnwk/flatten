package flatten

import (
	"log"
	"reflect"
	"strings"
)

// Struct flattens a struct with json tags
func Struct(s interface{}, m map[string]interface{}) {
	flatten("", s, m)
}

func flatten(key string, s interface{}, m map[string]interface{}) {
	val := reflect.ValueOf(s)

	// if its a pointer, resolve its value
	if val.Kind() == reflect.Ptr {
		val = reflect.Indirect(val)
	}

	// should double check we now have a struct (could still be anything)
	if val.Kind() != reflect.Struct {
		log.Fatal("unexpected type")
	}

	for i := 0; i < val.NumField(); i++ {
		tag := strings.SplitN(val.Type().Field(i).Tag.Get("json"), ",", 2)[0]
		if key != "" {
			tag = key + "." + tag
		}

		switch reflect.ValueOf(val.Field(i).Interface()).Kind() {
		case reflect.Struct:
			flatten(tag, val.Field(i).Interface(), m)
		default:
			if !reflect.DeepEqual(val.Field(i).Interface(), reflect.Zero(val.Field(i).Type()).Interface()) {
				m[tag] = val.Field(i).Interface()
			}
		}
	}
}
