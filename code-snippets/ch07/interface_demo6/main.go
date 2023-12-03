package main

import "fmt"

type Mover interface {
	Move()
}

type Dog struct {
	Name string
}

func (d *Dog) Move() {
	fmt.Println("狗在跑~")
}

type Car struct {
	Brand string
}

func (c *Car) Move() {
	fmt.Println("汽车在跑~")
}

func main() {
	var m Mover
	fmt.Println(m == nil)
	//m.Move() // panic: runtime error: invalid memory address or nil pointer dereference
	fmt.Printf("%T\n", m)

	m = &Dog{Name: "旺财"}
	fmt.Printf("%T\n", m) // *main.Dog

	m = new(Car)
	fmt.Printf("%T\n", m) // *main.Car
	fmt.Println(m == nil) // false
	m.Move()

	m = nil

	var (
		x Mover = new(Dog)
		y Mover = new(Car)
	)
	fmt.Println(x == y) // false

	//var z interface{} = []int{1, 2, 3}
	//fmt.Println(z == z) // panic: runtime error: comparing uncomparable type []int

	var n Mover = &Dog{Name: "旺财"}
	v, ok := n.(*Dog)
	if ok {
		fmt.Println("类型断言成功")
		v.Name = "富贵" // 变量v是*Dog类型
	} else {
		fmt.Println("类型断言失败")
	}
}

// justifyType 对传入的空接口类型变量x进行类型断言
func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}
