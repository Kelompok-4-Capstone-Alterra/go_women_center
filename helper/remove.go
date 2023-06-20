package helper

import (
	"reflect"
	"strings"
)

func RemoveWhiteSpace(r interface{}) {
	value := reflect.ValueOf(r)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
		// remove white space from string field
		for i := 0; i < value.NumField(); i++ {
			field := value.Field(i)
			if field.Kind() == reflect.String {
				field.SetString(strings.TrimSpace(field.String()))
			}
		}
	}
}