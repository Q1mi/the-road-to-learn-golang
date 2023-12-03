package main

import (
	"fmt"
	"sort"
)

// 3-3 映射

func main() {
	var m map[string]string
	fmt.Println(m == nil) // true

	// 创建一个key是string，value是int的映射
	m1 := map[string]int{
		"七米": 100,
		"张三": 60,
	}
	fmt.Println(m1) // map[七米:100 张三:60]

	m2 := make(map[string]string)
	fmt.Println(m2 == nil) // false

	scoreMap := make(map[string]int, 2)
	// 添加元素
	scoreMap["七米"] = 100
	scoreMap["张三"] = 60
	fmt.Printf("scoreMap:%v\n", scoreMap)
	fmt.Printf("type:%T\n", scoreMap)

	// 根据key访问元素
	fmt.Println(scoreMap["七米"])

	// 定义一个存储用户年龄的映射
	ageMap := map[string]int{
		"七米": 30,
		"张三": 28,
	}

	// 尝试查找小明的年龄
	name := "小明"
	fmt.Println(ageMap[name]) // 0

	// 如果key存在，则ok为true，v为对应的值
	// 如果key不存在，则ok为false，v为值类型的零值
	v, ok := ageMap[name]
	if ok {
		fmt.Printf("%s的年龄：%v\n", name, v)
	} else {
		fmt.Printf("%s的年龄：保密\n", name)
	}

}

func deleteDemo() {
	scoreMap := make(map[string]int, 4)
	scoreMap["七米"] = 100
	scoreMap["张三"] = 60
	scoreMap["小明"] = 85
	fmt.Println(scoreMap)  // map[七米:100 小明:85 张三:60]
	delete(scoreMap, "小明") // 将小明:85从映射中删除
	fmt.Println(scoreMap)  // map[七米:100 张三:60]
}

func rangeDemo() {
	scoreMap := make(map[string]int, 4)
	scoreMap["七米"] = 100
	scoreMap["张三"] = 60
	scoreMap["小明"] = 85
	for name, score := range scoreMap {
		fmt.Println(name, score)
	}

	// 只遍历取出映射中的key
	for name := range scoreMap {
		fmt.Println(name)
	}

	// 只遍历取出映射中的值
	for _, score := range scoreMap {
		fmt.Println(score)
	}
}

func sortRangeDemo() {
	var scoreMap = map[string]int{
		"stu03": 88,
		"stu01": 92,
		"stu05": 60,
		"stu02": 82,
		"stu06": 98,
		"stu04": 92,
	}

	// 取出映射中的所有key存入切片keys
	var keys = make([]string, 0, len(scoreMap))
	for key := range scoreMap {
		keys = append(keys, key)
	}
	// 对keys切片进行排序
	sort.Strings(keys)
	// 按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}

func rangeMapSlice() {
	var mapSlice = make([]map[string]string, 3)

	// 对切片中的映射元素进行初始化
	mapSlice[0] = make(map[string]string, 4)

	mapSlice[0]["name"] = "七米"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "北京"

	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
}

func rangeSliceMap() {
	// 声明一个键为字符串，值为字符串切片的映射
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		// 如果key不存在，则初始化一个值切片
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}
