package main

import (
	"errors"
	"fmt"
)

func modify(x int) {
	x = 100 // 此处改变形参的值不会影响实参
}

func modify2(y *int) {
	*y = 100 // 此处传入的是实参的地址，根据地址赋值会影响实参
}

func sayHello(name string) {
	fmt.Println("hello", name)
}

//func main() {
//	var f func(string) // 声明函数类型变量f
//	fmt.Println(f == nil)
//	f = sayHello // 将sayHello赋值给变量f
//	f("七米")      // 输出：hello 七米
//}

// add 两个整数相加
func add(x, y int) int {
	return x + y
}

// sub 两个整数相减
func sub(x, y int) int {
	return x - y
}

// calc 把两个整数传入给定op函数进行计算并返回结果
func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

//func main() {
//	addRes := calc(10, 20, add)
//	fmt.Println(addRes) // 30
//	subRes := calc(10, 20, sub)
//	fmt.Println(subRes) // -10
//}

// do 根据传入的符号返回对应的函数和可能出现的错误
func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}

func main() {
	// 函数有多个返回值时，使用两个变量分别接收对应的返回值
	f, err := do("+")
	if err != nil {
		// 如果出现错误就退出程序
		fmt.Println(err)
		return
	}
	res := f(10, 20)
	fmt.Println("res:", res) // 输出：res:30
}

//func main(){
//	x := 1
//	modify(x)
//	fmt.Println(x)  // 1
//
//	y := 1
//	modify2(&y)
//	fmt.Println(y)  // 100
//}
