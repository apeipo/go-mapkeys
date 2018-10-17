package gomod

import (
	"reflect"
)

// return map keys as interface
// panic if m is not map type
func MapKeys(m interface{}) interface{} {
	mtyp := reflect.TypeOf(m)
	v := reflect.ValueOf(m)
	keys := v.MapKeys() // panic if not map type
	rkeys := reflect.MakeSlice(reflect.SliceOf(mtyp.Key()), 0, v.Len())
	for _, key := range keys {
		rkeys = reflect.Append(rkeys, key)
	}
	return rkeys.Interface()
}
