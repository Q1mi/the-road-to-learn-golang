package main

import "fmt"

type Cat struct{}

func (c Cat) Say() {
	fmt.Println("喵喵喵~")
}

type Dog struct{}

func (d Dog) Say() {
	fmt.Println("汪汪汪~")
}

type Sheep struct{}

func (s Sheep) Say() {
	fmt.Println("咩咩咩~")
}

// MakeCatHungry 猫饿了会喵喵喵~
func MakeCatHungry(c Cat) {
	c.Say()
}

// MakeSheepHungry 羊饿了会咩咩咩~
func MakeSheepHungry(s Sheep) {
	s.Say()
}

// Sayer 必须提供一个Say方法
type Sayer interface {
	Say()
}

// MakeHungry 饿肚子了...
func MakeHungry(s Sayer) {
	s.Say()
}

func main() {
	c := Cat{}
	c.Say()
	d := Dog{}
	d.Say()

	var x Sayer // 声明一个Sayer类型的变量x
	a := Cat{}  // 声明一个Cat类型变量a
	b := Dog{}  // 声明一个Dog类型变量b
	x = a       // 可以把Cat类型变量直接赋值给x
	x.Say()     // 喵喵喵
	x = b       // 可以把Dog类型变量直接赋值给x
	x.Say()     // 汪汪汪
}
