package main

import "fmt"

// Mover 定义一个接口类型
type Mover interface {
	Move()
}

// Dog 狗结构体类型
type Dog struct{}

// Move 使用值接收者定义Move方法实现Mover接口
func (d Dog) Move() {
	fmt.Println("狗会动")
}

// Cat 猫结构体类型
type Cat struct{}

// Move 使用指针接收者定义Move方法实现Mover接口
func (c *Cat) Move() {
	fmt.Println("猫会动")
}

func main() {
	var x Mover    // 声明一个Mover类型的变量x
	var d1 = Dog{} // d1是Dog类型
	x = d1         // 可以将d1当成Mover类型
	x.Move()

	var d2 = &Dog{} // d2是Dog指针类型
	x = d2          // 也可以将d2当成Mover类型
	x.Move()

	var c1 = &Cat{} // c1是*Cat类型
	x = c1          // 可以将c1当成Mover类型
	x.Move()

	//// 下面的代码无法通过编译
	//var c2 = Cat{} // c2是Cat类型
	//x = c2         // 不能将c2当成Mover类型

}
