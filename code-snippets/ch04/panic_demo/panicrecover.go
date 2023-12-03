package main

import (
	"fmt"
	"time"
)

// 此程序将带着未被恢复的恐慌1而崩溃退出。

func demo1() {
	demo2()
	defer func() {
		recover()
	}()
	//panic("panic in demo2")
}

func demo2() {
	// ...
	panic("panic in demo2")
}

func demo3() {
	defer recover() // 不要使用defer直接调用recover
	panic("panic in demo1")
}

func demo4() {
	func() {
		defer recover() // 无效
	}()
	panic("panic in demo4")
}

func demo5() {
	func() {
		defer func() {
			recover() // 无效
		}()
	}()
	panic("panic in demo5")
}

func demo6() {
	defer func() {
		recover() // 空操作
	}()
	go func() {
		defer func() {
			recover()
		}()
		panic("panic in demo6")
	}()

	time.Sleep(time.Second * 5)
}

func main() {

	//demo1()
	//demo3()
	demo4()
	//demo5()
	//demo6()
}

func funcA() {
	fmt.Println("func A")
}

func funcB() {
	panic("panic in B")
}

func funcC() {
	fmt.Println("func C")
}

//func main() {

//}
