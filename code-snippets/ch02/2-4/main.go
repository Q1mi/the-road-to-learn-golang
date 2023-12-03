package main

import (
	"fmt"
	"math"
)

// 2-4 基本数据类型

var f1 float32 = 1.23 // float32
var f2 = 1.23         // float64

func main() {
	var c1 complex64
	c1 = 1 + 2i
	var c2 complex128
	c2 = 2 + 3i
	c3 := 4 + 5i // complex128

	c := complex(1.2, 3.4) // 构造复数c (1.2+3.4i)
	fc1 := real(c)         // 获取复数c的实部 1.2
	fc2 := imag(c)         // 获取复数c的虚部 3.4

	v1 := 0b00101101 // 代表二进制的 101101，相当于十进制的 45
	v2 := 0o377      // 代表八进制的 377，相当于十进制的 255
	v3 := 0x1p-2     // 代表十六进制的 1 除以 2²，也就是 0.25

	v4 := 123_456 // v4 := 123456

	var no bool // 默认为false
	ok := true  // true

	s1 := "hello"
	s2 := "你好" // 支持直接添加非ASCII码字符

	path := "c:\\Code\\lesson1\\app.exe"

	s3 := `第一行 
		第二行 
		第三行
		`
	fmt.Println(s3)

	var b1 = 'a'      // 字符变量默认类型为rune型
	var b2 byte = 'a' // byte型，本质上是uint8
	b3 := '中'         // rune型，本质上是int32

}

func changeString() {
	s1 := "big"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

	s2 := "白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红'
	fmt.Println(string(runeS2))
}

// 遍历字符串
func traversalString() {
	s := "hello中国"
	for i := 0; i < len(s); i++ { // byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { // rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}

func sqrtDemo() {
	var a, b = 3, 4
	var c int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
