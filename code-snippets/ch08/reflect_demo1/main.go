package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a float32 = 3.14
	ta := reflect.TypeOf(a)
	fmt.Printf("type:%v\n", ta) // type:float32
	var b int64 = 100
	tb := reflect.TypeOf(b)
	fmt.Printf("type:%v\n", tb) // type:int64
}
