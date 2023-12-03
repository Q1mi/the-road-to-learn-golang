package main

import (
	"fmt"
)

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	x := 1
	y := 2
	defer func() {
		r := calc("A", x, y)
		calc("AA", x, r)
	}()
	x = 10
	defer func() {
		r := calc("B", x, y)
		calc("BB", x, r)
	}()
	y = 20

}
