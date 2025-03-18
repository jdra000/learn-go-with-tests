package main

import "reflect"

// * interface{} means any type. You lose type safety with it
// * As a writer of such a function, you have to be able to inspect anything
// * that has been passed to you and try and figure out what the type is and what you can do with it.
// * This is done using reflection.
func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String()) // * Calls the function passed as a parameter with the string value
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ { // * Iterate over the fields
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ { // * Iterate over the indexes
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() { // * Iterate over the keys
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan:
		for {
			if v, ok := val.Recv(); ok { // * Recv() allows to iterate through all values sent through channel until it was closed
				walkValue(v)
			} else {
				break
			}
		}
	case reflect.Func:
		response := val.Call(nil)
		for _, res := range response {
			walkValue(res)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x) // * ValueOf returns the value of a given variable

	if val.Kind() == reflect.Pointer {
		val = val.Elem() // * Dereference the pointer if that's the case
	}
	return val
}
