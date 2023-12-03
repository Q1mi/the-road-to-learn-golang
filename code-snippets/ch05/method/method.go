package main

import "fmt"

// 方法与接收者

//// Person 人
//type Person struct {
//	name string
//	age  int8
//}
//
//// Dream 做梦
//func (p Person) Dream() {
//	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
//}

// MyInt 基于内置int定义的类型
type MyInt int

// greaterThan100 判断当前整数是否大于100的方法
func (m MyInt) greaterThan100() bool {
	return m > 100
}

// StringSet 基于映射自定义一个字符串集合类型
type StringSet map[string]bool

// Has 判断集合中是否包含指定key的方法
func (ss StringSet) Has(key string) bool {
	return ss[key]
}

// Add 向集合中添加指定key的方法
func (ss StringSet) Add(key string) {
	ss[key] = true
}

// Remove 从集合中移除指定key的方法
func (ss StringSet) Remove(key string) {
	delete(ss, key)
}

// FilterFunc 基于函数自定义一个类型
type FilterFunc func(int)

// Do 执行函数的方法
func (f FilterFunc) Do(n int) {
	f(n)
}

// Employee 职员
type Employee struct {
	name   string // 姓名
	salary int    // 薪资
}

// SayHi 打招呼
func (e Employee) SayHi() {
	fmt.Printf("你好！ 我是%s。\n", e.name)
}

// raise 涨薪
func (e *Employee) raise(n int) {
	e.salary += n
}

// setSalary 设置薪水数
func (e Employee) setSalary(n int) {
	e.salary = n
	fmt.Println("setSalary:", e.salary)
}

//func (i int) method() {
//	// 不能为其他包的类型定义方法
//}
//
//func (i []int) method() {
//	// 不可为非定义的类型命名
//}
//
//// PE 属于基于指针定义的类型
//type PE *Employee
//
//func (p PE) method() {
//	// 不能为指针类型定义方法
//}

//type P2 = []int
//
//func (p P2) M() {
//	fmt.Println("p2.M")
//}

func main() {

	// 创建结构体变量e1
	e1 := Employee{
		name:   "张三",
		salary: 25,
	}

	// 调用方法
	e1.SayHi()

	fmt.Println(e1.salary) // 25
	e1.setSalary(10000)    // setSalary: 10000
	fmt.Println(e1.salary) // 25

	// 调用涨薪方法
	e1.raise(10)           // 等价于(&e1).raise(10)
	fmt.Println(e1.salary) // 35

	// 创建结构体指针变量e2
	e2 := &Employee{
		name:   "小红",
		salary: 20,
	}

	// 调用打招呼方法
	e2.SayHi() // 等价于(*e2).SayHi()
}
