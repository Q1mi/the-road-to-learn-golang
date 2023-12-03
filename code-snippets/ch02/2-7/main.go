package main

import "fmt"

// 2-7 流程控制语句

func ifDemo1() {
	score := 65
	if score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}

func ifDemo2() {
	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i < 5; i++ {
		if i == 3 {
			// 当 i== 3时跳过本次循环，直接进入下一次循环
			continue
		}
		fmt.Println(i)
	}

	for i := 0; i < 5; i++ {
		if i == 3 {
			// 当 i== 3时跳出 for 循环
			break
		}
		fmt.Println(i)
	}

	var finger int
	if finger == 1 {
		fmt.Println("大拇指")
	} else if finger == 2 {
		fmt.Println("食指")
	} else if finger == 3 {
		fmt.Println("中指")
	} else if finger == 4 {
		fmt.Println("无名指")
	} else if finger == 5 {
		fmt.Println("小拇指")
	} else {
		fmt.Println("无效的数字")
	}

	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的数字")
	}

	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}

	score := 92
	switch {
	case score >= 95:
		fmt.Println("S")
	case score >= 90 && score < 95:
		fmt.Println("A")
	case score >= 60:
		fmt.Println("B")
	default:
		fmt.Println("C")
	}

	s := "c"
	switch {
	case s == "a":
		fmt.Println("a")
	case s == "b":
		fmt.Println("b")
		fallthrough // 继续执行下一个分支
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}

	var breakFlag bool
	for i := 0; i < 5; i++ {
		for j := 10; j < 15; j++ {
			if j == 12 {
				// 设置退出标签
				breakFlag = true
				break
			}
			fmt.Println(i, j)
		}
		// 外层for循环判断
		if breakFlag {
			fmt.Println("结束for循环")
			break
		}
	}

	for i := 0; i < 5; i++ {
		for j := 10; j < 15; j++ {
			if j == 12 {
				// 跳转到退出标签
				goto breakTag
			}
			fmt.Println(i, j)
		}
	}
	return
	// 标签
breakTag:
	fmt.Println("结束for循环")

}

func breakDemo1() {
breakLabel:
	for i := 0; i < 5; i++ {
		for j := 10; j < 15; j++ {
			if j == 2 {
				break breakLabel
			}
			fmt.Println(i, j)
		}
	}
	fmt.Println("循环结束")
}

func continueDemo() {
forLoop:
	for i := 0; i < 5; i++ {
		for j := 10; j < 15; j++ {
			if i == 2 && j == 12 {
				continue forLoop
			}
			fmt.Println(i, j)
		}
	}
}
