package main

import (
	"fmt"
	"regexp"
)

func div(x, y int) int {
	if y == 0 {
		panic("0不能做除数")
	}
	return x / y
}

var phoneRE = regexp.MustCompile(`^1[3-9]\d{9}$`)

func main() {
	fmt.Println(123)
	panic(456) // 手动panic,程序异常退出，不会执行后续语句
	fmt.Println(789)
}
