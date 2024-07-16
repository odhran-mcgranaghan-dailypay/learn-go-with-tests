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

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			val := val.MapIndex(key)
			walk(val.Interface(), fn)
		}
	case reflect.Chan:
		// iterate over the values sent to the channel
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walk(v.Interface(), fn)	
		}
	case reflect.Func:
		valueFunctionResult := val.Call(nil)
		for _, res := range valueFunctionResult {
			walkValue(res)
		}
	}
}
