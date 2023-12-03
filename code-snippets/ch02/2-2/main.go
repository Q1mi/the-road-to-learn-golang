package main

import (
	"fmt"
)

var name = "七米" // string型
var age = 18    // int型

// age := 18 // 此处无法使用简短变量声明

func foo() (int, string) {
	return 18, "七米"
}

func main() {
	age := 20 // 函数中可以使用简短变量声明
	fmt.Println(age)

	x, _ := foo()
	_, y := foo()
	fmt.Println("x=", x)
	fmt.Println("y=", y)
}
