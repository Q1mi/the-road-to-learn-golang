package main

import "fmt"

// 3-1 数组

func main() {
	// 数组
	/*
		var a [3]int

		var b [5]string

		var n = 10
		// var c = [n]int // 数据的长度n为变量，导致编译失败

		var x [3]int
		var y [4]int
		// x = y // 不可以这样做，因为此时x和y是不同的类型

		var x [3]int
		fmt.Println(x[0]) // 输出数组的第一个元素
		// fmt.Println(x[-1]) // 不支持使用负数作为索引

		fmt.Println(len(x))      // 3
		fmt.Println(x[len(x)-1]) // 输出数组x的最后一个元素
	*/
	var s1 [3]int
	fmt.Println(s1) // [0 0 0]

	var s2 [2]string
	fmt.Println(s2)         // [ ]，空字符串看不出效果，可使用下面的语句查看
	fmt.Printf("%#v\n", s2) // [2]string{"", ""}

	/* 方法一
	var testArray [3]int        // 数组元素会初始化为int型的零值
	var numArray = [3]int{1, 2} // 使用指定的初始值完成初始化
	var cityArray = [4]string{"北京", "上海", "广州", "深圳"}
	fmt.Println(testArray) // [0 0 0]
	fmt.Println(numArray)  // [1 2 0]
	fmt.Println(cityArray) // [北京 上海 广州 深圳]
	*/

	// 方法二
	var testArray [3]int
	var numArray = [...]int{1, 2}
	var cityArray = [...]string{"北京", "上海", "广州", "深圳"}
	fmt.Println(testArray)                          // [0 0 0]
	fmt.Println(numArray)                           // [1 2]
	fmt.Printf("type of numArray:%T\n", numArray)   // type of numArray:[2]int
	fmt.Println(cityArray)                          // [北京 上海 广州 深圳]
	fmt.Printf("type of cityArray:%T\n", cityArray) // type of cityArray:[4]string

	/* 方法三
	a := [...]int{1: 1, 3: 5}       // 索引1处的值为1，索引3处的值为5
	fmt.Println(a)                  // [0 1 0 5]
	fmt.Printf("type of a:%T\n", a) // type of a:[4]int
	*/

	// 方法1：for循环遍历
	for i := 0; i < len(cityArray); i++ {
		fmt.Println(cityArray[i])
	}

	// 方法2：for range遍历
	// 输出索引和元素值
	for idx, city := range cityArray {
		fmt.Println(idx, city)
	}

	// 方法2：for range遍历
	// 只输出元素值
	for _, city := range cityArray {
		fmt.Println(city)
	}

	cityArray2 := [2][4]string{
		{"北京", "上海", "广州", "深圳"},
		{"杭州", "成都", "武汉", "重庆"},
	}
	fmt.Println(cityArray2)       // [[北京 上海 广州 深圳] [杭州 成都 武汉 重庆]]
	fmt.Println(cityArray2[1][2]) // 支持索引取值:武汉

	for _, v1 := range cityArray2 { // 获取内层数组
		for _, city := range v1 { // 遍历内层数组
			fmt.Printf("%s\t", city)
		}
		fmt.Println()
	}

	// 多维数组第一层支持使用...推导数组长度的写法
	a := [...][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(a)
	// 错误的写法
	// 多维数组的内层不支持使用...符号
	// b := [3][...]string{
	// 	{"北京", "上海"},
	// 	{"广州", "深圳"},
	// 	{"成都", "重庆"},
	// }

	// 数组比较
	// var a [3]int
	// var b = [...]int{1, 2, 3}
	// var c = [3]int{1, 2, 3}
	// fmt.Println(a == b, b == c, a == c) // false true false

}
