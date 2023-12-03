package main

import "fmt"

// 练习题

func structEx1() {
	type student struct {
		name string
		age  int
	}

	m := make(map[string]*student, 4)
	s := []student{
		{name: "七米", age: 18},
		{name: "张三", age: 23},
		{name: "小红", age: 20},
	}

	for i := range s {
		m[s[i].name] = &s[i]
	}

	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}

// Person 人
type Person struct {
	name   string
	age    int8
	dreams []string
}

// SetDreams 设置梦想的方法
func (p *Person) SetDreams(dreams []string) {
	p.dreams = dreams
}

func main() {
	structEx1()

	//p1 := Person{name: "小王子", age: 18}
	//// 定义一个切片变量data
	//data := []string{"吃饭", "睡觉", "打豆豆"}
	//// 将切片变量传入SetDreams方法
	//p1.SetDreams(data)
	//
	//// 在修改切片变量data
	//data[1] = "不睡觉"
	//fmt.Println(p1.dreams) // [吃饭 不睡觉 打豆豆]
}
