package reflection

import (
	"reflect"
)

func getValue(x interface{}) reflect.Value {
	value := reflect.ValueOf(x)
	// handle value passed as pointer
	if value.Kind() == reflect.Ptr {
		// Elem returns the value that the interface v contains
		// or that the pointer v points to.
		value = value.Elem()
	}
	return value
}

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	if val.Kind() == reflect.Slice {
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), fn)
		}
		return
	}

	// iterate through fields of x
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}
}
