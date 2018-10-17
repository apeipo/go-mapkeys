package maputil

import (
	"reflect"
	"fmt"
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

func SafeMapKeys(m interface{}) (r interface{}, err error) {
	defer func() {
		if p := recover(); p != nil {
			r = nil
			err = fmt.Errorf("%s", p)
		}
	}()
	mtyp := reflect.TypeOf(m)
	v := reflect.ValueOf(m)
	keys := v.MapKeys() // panic if not map type
	rkeys := reflect.MakeSlice(reflect.SliceOf(mtyp.Key()), 0, v.Len())
	for _, key := range keys {
		rkeys = reflect.Append(rkeys, key)
	}
	return rkeys.Interface(), nil
}
