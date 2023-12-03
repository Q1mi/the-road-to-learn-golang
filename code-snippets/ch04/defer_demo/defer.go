package main

import "fmt"

func f1() int {
	x := 5
	defer func() {
		x++ // 只修改x的值，不会更新返回值
	}()
	return x // 返回值=5
}

func f2() (x int) {
	// 已提前声明好返回值变量x
	defer func() {
		x++ // 会更新返回值
	}()
	return 5 // 返回值x=5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++ // 只修改x的值，不会更新返回值
	}()
	return x // 返回值y=5
}
func f4() (x int) {
	defer func(x int) {
		x++ // 修改的是匿名函数局部变量x，不会更新返回值
	}(x)
	return 5 // 返回值x=5
}

func main() {
	fmt.Println(f1()) // 5
	fmt.Println(f2()) // 6
	fmt.Println(f3()) // 5
	fmt.Println(f4()) // 5
}
