package main

import "fmt"

// 内置函数：new和make

// newAndMake new和make示例
func newAndMake() {
	var a *int // nil
	*a = 100

	var s []int // nil
	s[0] = 100

	var m map[string]int // nil
	m["七米"] = 100

	//fmt.Println(a, s, m)
}

func newDemo1() {
	a := new(int)
	b := new(bool)
	fmt.Printf("%T\n", a) // *int
	fmt.Printf("%T\n", b) // *bool
	fmt.Println(*a)       // 0
	fmt.Println(*b)       // false
}

func newDemo2() {
	var a *int
	a = new(int)
	*a = 10
	fmt.Println(*a)
}

// makeDemo make初始化
func makeDemo() {

	var s []int // nil
	s = make([]int, 1)
	s[0] = 100

	var m map[string]int // nil
	m = make(map[string]int, 2)
	m["七米"] = 100

	//fmt.Println(a, s, m)
}

func main() {
	newAndMake()
}
