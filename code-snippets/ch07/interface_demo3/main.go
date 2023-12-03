package main

import "fmt"

// Sayer 接口
type Sayer interface {
	Say()
}

// Mover 接口
type Mover interface {
	Move()
}

type Dog struct {
	Name string
}

// 实现Sayer接口
func (d Dog) Say() {
	fmt.Printf("%s会叫汪汪汪\n", d.Name)
}

// 实现Mover接口
func (d Dog) Move() {
	fmt.Printf("%s会动\n", d.Name)
}

// Car 汽车结构体类型
type Car struct {
	Brand string
}

// Move Car类型实现Mover接口
func (c Car) Move() {
	fmt.Printf("%s速度70迈\n", c.Brand)
}

func main() {
	var d = Dog{Name: "旺财"}

	var s Sayer = d
	var m Mover = d

	s.Say()  // 对Sayer类型调用Say方法
	m.Move() // 对Mover类型调用Move方法

	var obj Mover
	obj = Dog{Name: "旺财"}
	obj.Move()

	obj = Car{Brand: "宝马"}
	obj.Move()

}
