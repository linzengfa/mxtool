// Copyright (c) 2020.
// ALL Rights reserved.
// @Description mxconv.go
// @Author moxiao
// @Date 2020/11/22 10:19

package mxconv

import (
	"fmt"
	"reflect"
)

//结构体转Map-带tag
func StructToMapWithTag(obj interface{}, tag string) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	fmt.Println(t.String())
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		if len(tag) > 0 && len(t.Field(i).Tag.Get(tag)) > 0 {
			data[t.Field(i).Tag.Get(tag)] = v.Field(i).Interface()
		} else {
			data[t.Field(i).Name] = v.Field(i).Interface()
		}
	}
	return data
}

//结构体转Map
func StructToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}
