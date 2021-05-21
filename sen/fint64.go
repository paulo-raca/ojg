// Copyright (c) 2021, Peter Ohler, All rights reserved.

package sen

import (
	"reflect"
	"strconv"
	"unsafe"
)

var int64AppendFuncs = [8]appendFunc{
	appendInt64,
	appendInt64AsString,
	appendInt64NotEmpty,
	appendInt64NotEmptyAsString,
	iappendInt64,
	iappendInt64AsString,
	iappendInt64NotEmpty,
	iappendInt64NotEmptyAsString,
}

func appendInt64(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, interface{}, bool, bool) {
	buf = append(buf, fi.jkey...)
	buf = strconv.AppendInt(buf, *(*int64)(unsafe.Pointer(addr + fi.offset)), 10)

	return buf, nil, true, false
}

func appendInt64AsString(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, interface{}, bool, bool) {
	buf = append(buf, fi.jkey...)
	buf = append(buf, '"')
	buf = strconv.AppendInt(buf, *(*int64)(unsafe.Pointer(addr + fi.offset)), 10)
	buf = append(buf, '"')

	return buf, nil, true, false
}

func appendInt64NotEmpty(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, interface{}, bool, bool) {
	v := *(*int64)(unsafe.Pointer(addr + fi.offset))
	if v == 0 {
		return buf, nil, false, false
	}
	buf = append(buf, fi.jkey...)
	buf = strconv.AppendInt(buf, v, 10)

	return buf, nil, true, false
}

func appendInt64NotEmptyAsString(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, interface{}, bool, bool) {
	v := *(*int64)(unsafe.Pointer(addr + fi.offset))
	if v == 0 {
		return buf, nil, false, false
	}
	buf = append(buf, fi.jkey...)
	buf = append(buf, '"')
	buf = strconv.AppendInt(buf, v, 10)
	buf = append(buf, '"')

	return buf, nil, true, false
}

func iappendInt64(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, interface{}, bool, bool) {
	buf = append(buf, fi.jkey...)
	buf = strconv.AppendInt(buf, rv.FieldByIndex(fi.index).Interface().(int64), 10)

	return buf, nil, true, false
}

func iappendInt64AsString(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, interface{}, bool, bool) {
	buf = append(buf, fi.jkey...)
	buf = append(buf, '"')
	buf = strconv.AppendInt(buf, rv.FieldByIndex(fi.index).Interface().(int64), 10)
	buf = append(buf, '"')

	return buf, nil, true, false
}

func iappendInt64NotEmpty(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, interface{}, bool, bool) {
	v := rv.FieldByIndex(fi.index).Interface().(int64)
	if v == 0 {
		return buf, nil, false, false
	}
	buf = append(buf, fi.jkey...)
	buf = strconv.AppendInt(buf, v, 10)

	return buf, nil, true, false
}

func iappendInt64NotEmptyAsString(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, interface{}, bool, bool) {
	v := rv.FieldByIndex(fi.index).Interface().(int64)
	if v == 0 {
		return buf, nil, false, false
	}
	buf = append(buf, fi.jkey...)
	buf = append(buf, '"')
	buf = strconv.AppendInt(buf, v, 10)
	buf = append(buf, '"')

	return buf, nil, true, false
}