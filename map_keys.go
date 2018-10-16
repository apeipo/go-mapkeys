package gomod

import (
	"reflect"
)

// return map keys as interface
// panic if m is not map type
func MapKeys(m interface{}) interface{} {
	mtyp := reflect.TypeOf(m)
	if mtyp.Kind() != reflect.Map {
		panic("not map type in map keys call")
	}
	v := reflect.ValueOf(m)
	keys := v.MapKeys()
	rkeys := reflect.MakeSlice(reflect.SliceOf(mtyp.Key()), 0, v.Len())
	for _, key := range keys {
		rkeys = reflect.Append(rkeys, key)
	}
	return rkeys.Interface()
}
