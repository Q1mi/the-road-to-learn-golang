package main

import "fmt"

// 类型声明

type MyInt int // 基于int8类型声明一个MyInt类型

type Roster []string // 基于字符串切片声明一个Roster类型

type report map[string]int // 基于映射声明一个成绩单类型

type Notify func(string, report) // 基于函数申明一个Notify类型

func typeDeclare() {
	type MyInt int

	var a int = 255
	var b MyInt

	// b = a // 不能把int类型的a直接赋值给MyInt类型的b
	b = MyInt(a)   // 显式类型转换
	fmt.Println(b) // 255
}

func typeAlias() {

	// 类型别名
	type MyInt = int

	var a = 255
	var b MyInt

	b = a          // int类型的a可以直接赋值给MyInt类型的b
	fmt.Println(b) // 255
}
func main() {

}
