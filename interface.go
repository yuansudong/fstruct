package fstruct

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
)

// FillFromSlice 从数据源为slice的切片中填充对象
// func FillFromSlice(
// 	iObject interface{}, // iObject 需要被填充的对象
// 	iDataSource map[string]interface{}, // iDataSource 数据源
// 	iTag string, //  对象中对应的key
// ) error {
// 	return nil
// }

// FillFromMap 从map中填充
func FillFromMap(
	io interface{},
	ids map[string]interface{},
	itag string,
) error {
	if io == nil {
		return ErrIOIsNil
	}
	pv := reflect.ValueOf(io)
	if pv.Kind() != reflect.Ptr {
		return ErrNoPtr
	}
	pv = pv.Elem()
	bt := reflect.TypeOf(io).Elem()
	if pv.Kind() != reflect.Struct {
		return ErrNoStruct
	}
	iFieldNum := bt.NumField()
	log.Println(" type num: ", bt.NumField(), " value num: ", iFieldNum)
	for index := 0; index < iFieldNum; index++ {
		tvv := bt.Field(index)
		tag := tvv.Tag.Get(itag)
		if tag == _Ignore {
			continue
		}
		tv := pv.Field(index)
		if !tv.CanSet() {
			continue
		}
		iVal, ok := ids[tag]
		if !ok {
			continue
		}
		sVal := fmt.Sprint(iVal)
		switch tv.Kind() {
		case reflect.String:
			tv.SetString(sVal)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			iv, err := strconv.ParseInt(sVal, 10, 64)
			if err != nil {
				return err
			}
			tv.SetInt(iv)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			iv, err := strconv.ParseUint(sVal, 10, 64)
			if err != nil {
				return err
			}
			tv.SetUint(iv)
		case reflect.Float32, reflect.Float64:
			iv, err := strconv.ParseFloat(sVal, 64)
			if err != nil {
				return err
			}
			tv.SetFloat(iv)
		default:
			return fmt.Errorf(_LayoutNoSupportType, tv.Kind().String(), tag)
		}
	}
	return nil
}
