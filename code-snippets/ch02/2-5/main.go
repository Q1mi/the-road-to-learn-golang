package main

import "fmt"

// 2-5 指针

func main() {
	v := 10
	p := &v         // 取变量v的地址，保存到指针p中
	fmt.Println(*p) // 10
	*p = 100        // 相当于 v = 100
	fmt.Println(v)  // 100

}
