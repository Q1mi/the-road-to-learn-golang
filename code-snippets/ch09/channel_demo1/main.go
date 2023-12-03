package main

import "fmt"

// 通道类型示例

func f() {

	var ch1 chan int
	var ch2 = make(chan int)

	fmt.Println(ch1 == nil)
	fmt.Println(ch2 == nil)
}

func f2(ch chan int) {
	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("通道已关闭")
			break
		}
		fmt.Printf("v:%#v ok:%#v\n", v, ok)
	}
}

func f3(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)
	//f2(ch)
	f3(ch)
}
