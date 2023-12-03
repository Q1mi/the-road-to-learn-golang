package main

import (
	"fmt"
)

// Mover 接口
type Mover interface {
	Move()
}

// Animal 动物结构体
type Animal struct {
	Name string
	Mover
}

func (a Animal) Move() {
	fmt.Printf("%s会动~", a.Name)
}

func main() {

	var m Mover = Animal{Name: "test"}
	fmt.Println(m.(Animal).Name)

	//var a = Animal{Name: "马"}
	//a.Mover = a
	//a.Mover.Move()

}
