package utils

import "reflect"

func CopyStruct(s any) any {
	clone := reflect.New(reflect.TypeOf(s)).Elem()
	clone.Set(reflect.ValueOf(s))
	res := clone.Interface()
	return res
}
