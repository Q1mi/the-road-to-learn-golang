package main

import "fmt"

// 3-2 切片

func main() {
	var s1 []string // 声明一个字符串切片
	var s2 []int    // 声明一个整型切片

	fmt.Println(s1, len(s1), cap(s1)) // [] 0 0
	fmt.Println(s2, len(s2), cap(s2)) // [] 0 0

	// 对数组取切片
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3] // s := a[low:high]
	fmt.Printf("s:%v type:%T len:%v cap:%v\n", s, s, len(s), cap(s))

	// 对字符串取切片得到的还是字符串类型
	b := "hello world"
	s3 := b[1:3] // s3 := b[low:high]
	fmt.Printf("s2:%v type:%T len:%v\n", s3, s3, len(s3))

}

func appendDemo() {
	var s []int
	s = append(s, 1)       // 追加1个元素
	fmt.Println(s)         // [1]
	s = append(s, 2, 3, 4) // 追加多个元素
	fmt.Println(s)         // [1 2 3 4]
}

func appendDemo2() {
	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n",
			numSlice, len(numSlice), cap(numSlice), numSlice)
	}
}

func copyDemo() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, 5, 5)
	copy(s2, s1)    // 将切片s1中的元素复制到切片s2中
	fmt.Println(s1) // [1 2 3 4 5]
	fmt.Println(s2) // [1 2 3 4 5]
	s2[0] = 1000    // 修改s2
	fmt.Println(s1) // [1 2 3 4 5]
	fmt.Println(s2) // [1000 2 3 4 5]
}
