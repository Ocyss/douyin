package utils

import (
	"errors"
	"reflect"
)

func Merge(dst, src any) error {
	srcT, srcV := reflect.TypeOf(src), reflect.ValueOf(src)
	if srcT.Kind() == reflect.Ptr {
		srcT, srcV = srcT.Elem(), srcV.Elem()
	}
	if srcT.Kind() != reflect.Struct {
		return errors.New("仅支持 Struct 进行合并")
	}
	dstT, dstV := reflect.TypeOf(dst), reflect.ValueOf(dst)
	if dstT.Kind() != reflect.Ptr || dstT.Elem().Kind() != reflect.Struct {
		return errors.New("dst 必须为 Struct指针")
	} else {
		dstT, dstV = dstT.Elem(), dstV.Elem()
	}
	for i := 0; i < srcV.NumField(); i++ {
		curT, curV := srcT.Field(i), srcV.Field(i)
		f := dstV.FieldByName(curT.Name)
		if curV.Type() == f.Type() {
			f.Set(curV)
		}
	}
	return nil
}
