package snow

import (
	"answer/calc" // 导入同一个项目下的calc包
	"fmt"
)

// addDemo
func addDemo() {
	a, b := 10, 20
	ret := calc.Add(a, b) // 调用calc包的Add方法
	fmt.Println(ret)
}

// subDemo
func subDemo() {
	a, b := 10, 20
	ret := calc.Sub(a, b) // 调用calc包的Sub方法
	fmt.Println(ret)
}

// mulDemo
func mulDemo() {
	a, b := 10, 20
	ret := calc.Mul(a, b) // 调用calc包的Mul方法
	fmt.Println(ret)
}

// divDemo
func divDemo() {
	a, b := 10, 20
	ret := calc.Div(a, b) // 调用calc包的Div方法
	fmt.Println(ret)
}
