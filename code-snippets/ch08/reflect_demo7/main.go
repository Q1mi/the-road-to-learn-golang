package main

import (
	"fmt"
	"reflect"
)

// 反射三大定律

func main() {

	var x int64 = 10
	//v := reflect.ValueOf(x)
	//v.SetInt(100) // panic: reflect: reflect.Value.SetInt using unaddressable value

	v := reflect.ValueOf(&x)
	// 指针类型反射对象需要调用Elem方法得到原值
	v.Elem().SetInt(100)
	fmt.Println(v.Elem().Interface()) // 100
	fmt.Println(x)                    // 100
}
