package main

import (
	"fmt"
	"strings"
)

func adder() func(int) int {
	var x int
	// 返回一个匿名函数
	return func(y int) int {
		// 将y的值累加到其外层函数的局部变量x
		x += y
		return x
	}
}

// adder2 将x设置为参数更灵活，可以按需传入
func adder2(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

// makeSuffixFunc 返回一个给文件名添加指定后缀的函数
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		// 判断文件名是否包含指定后缀
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	var f = adder()
	fmt.Println(f(10)) // 10
	fmt.Println(f(20)) // 30
	fmt.Println(f(30)) // 60

	f1 := adder()
	fmt.Println(f1(40)) // 40
	fmt.Println(f1(50)) // 90

	// 创建一个添加 .jpg 后缀的函数
	jpgFunc := makeSuffixFunc(".jpg")
	// 创建一个添加 .txt 后缀的函数
	txtFunc := makeSuffixFunc(".txt")
	// 给文件名avatar添加.jpg后缀
	fmt.Println(jpgFunc("avatar")) // avatar.jpg
	// 给文件名readme添加.txt后缀
	fmt.Println(txtFunc("readme")) // readme.txt
}
