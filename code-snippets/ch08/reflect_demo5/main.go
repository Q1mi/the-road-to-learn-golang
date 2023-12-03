package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 空指针
	var a *int
	va := reflect.ValueOf(a)
	fmt.Println(va.IsNil())   // true
	fmt.Println(va.IsValid()) // true

	// nil值
	v := reflect.ValueOf(nil)
	fmt.Println(v.IsValid()) // false
	//fmt.Println(v.IsNil())   // panic

	// 匿名结构体变量
	b := struct{}{}
	// 尝试从结构体中查找"abc"字段
	vbf := reflect.ValueOf(b).FieldByName("abc")
	fmt.Println(vbf.IsValid()) // false
	//fmt.Println(vbf.IsNil())   // panic

	// 尝试从结构体中查找"abc"方法
	vbm := reflect.ValueOf(b).MethodByName("abc")
	fmt.Println(vbm.IsValid()) // false
	//fmt.Println(vbm.IsNil())   // panic
}
