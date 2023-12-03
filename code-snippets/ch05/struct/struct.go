package main

import (
	"fmt"
	"unsafe"
)

// 结构体相关

// Info 联系方式信息
type Info struct {
	Email string
	Phone string
}

// Student 学生
type Student struct {
	// 为该字段设置标签，
	// 指定JSON数据中的字段为stu_id
	ID int `json:"stu_id"` // 学号
	// 不为字段指定标签
	// 则JSON数据中默认使用字段名
	Name   string // 姓名
	Gender string // 性别
	Age    int    // 年龄
	// 结构体Info作为字段类型
	Contact Info // 联系方式
}

type Book struct {
	Title string `key1:"value1" key2:"value2"`
}

// node 单向链表的节点
type node struct {
	value int
	next  *node
}

// tree 树节点
type tree struct {
	value       int
	left, right *tree
}

// Student2 学生
type Student2 struct {
	ID, Age      int
	Name, Gender string
}

func structDemo1() {
	var s Student
	fmt.Printf("%#v\n", s) // main.Student{ID:0, Name:"", Gender:"", Age:0}
	fmt.Println(s.ID)      // 0
	fmt.Println(s.Name)    //
	fmt.Println(s.Gender)  //
	fmt.Println(s.Age)     // 0

	s.ID = 1
	s.Name = "七米"
	s.Gender = "男"
	s.Age = 18
	fmt.Printf("%#v\n", s) // main.Student{ID:1, Name:"七米", Gender:"男", Age:18}
}

func structDemo2() {
	var s Student

	s = Student{
		ID:     1,
		Name:   "七米",
		Gender: "男",
		Age:    18,
	}

	fmt.Printf("%#v\n", s) // main.Student{ID:1, Name:"七米", Gender:"男", Age:18}
}

func structDemo3() {
	s := Student{
		ID:     1,
		Name:   "七米",
		Gender: "男",
		Age:    18,
	}
	fmt.Printf("%#v\n", s) // main.Student{ID:1, Name:"七米", Gender:"男", Age:18}
	s2 := Student{
		ID:   2,
		Name: "张三",
	}
	fmt.Println(s2)
}

func structDemo4() {
	// Point 二维坐标
	type Point struct {
		x, y int
	}
	p := Point{10, 4}
	fmt.Println(p.x, p.y) // 10 4
}

func structDemo5() {

	s := Student{
		ID:     3,
		Name:   "小红",
		Gender: "女",
		Age:    18} // 可省略 ,
	fmt.Println(s) // {3 小红 女 18}
}

func structDemo6() {
	// Entry 学科和成绩记录
	type Entry struct {
		string
		*int
	}

	x := new(int)
	*x = 100

	a := Entry{string: "语文", int: x}
	fmt.Printf("%#v\n", a) // main.Entry{string:"语文", int:100}

	//var job Job
	//job.string = "学编程"

	fmt.Println(a)
}

// structDemo7 匿名结构体
func structDemo7() {
	tmp := struct {
		ID   int
		Info string
	}{
		ID:   123,
		Info: "just a test",
	}
	fmt.Println(tmp) // {123 just a test}
	fmt.Printf("%#v\n", tmp)
}

// structDemo8 结构体大小
func structDemo8() {
	// Test1 一个测试用结构体
	type Test1 struct {
		a int8 // 1byte
		b int8 // 1byte
	}

	var v Test1
	// unsafe.Sizeof函数返回变量的字节大小
	fmt.Println(unsafe.Sizeof(v)) // 2
}

// structDemo9 结构体大小
func structDemo9() {
	// Test2 一个测试用结构体
	type Test2 struct {
		a int8  // 1byte
		b int16 // 2bytes
		c int8  // 1byte
	}

	var v2 Test2
	// unsafe.Sizeof函数返回变量的字节大小
	fmt.Println(unsafe.Sizeof(v2)) // 6
	// unsafe.Alignof函数返回对齐要求
	fmt.Println(unsafe.Alignof(v2)) // 2

	type Test3 struct {
		a int8  // 1byte
		c int8  // 1byte
		b int16 // 2bytes
	}

	var v3 Test3
	// unsafe.Sizeof函数返回变量的字节大小
	fmt.Println(unsafe.Sizeof(v3)) // 4
	// unsafe.Alignof函数返回对齐要求
	fmt.Println(unsafe.Alignof(v3)) // 2
}

// structDemo10 结构体大小
func structDemo10() {
	// Test4 一个测试用结构体
	type Test4 struct {
		a bool  // 1byte
		b int64 // 8bytes
		c byte  // 1byte
	}

	var v4 Test4
	// unsafe.Sizeof函数返回变量的字节大小
	fmt.Println(unsafe.Sizeof(v4)) // 24
	// unsafe.Alignof函数返回对齐要求
	fmt.Println(unsafe.Alignof(v4)) // 8

	type Test5 struct {
		a bool  // 1byte
		c byte  // 1byte
		b int64 // 8bytes
	}

	var v5 Test5
	// unsafe.Sizeof函数返回变量的字节大小
	fmt.Println(unsafe.Sizeof(v5)) // 16
	// unsafe.Alignof函数返回对齐要求
	fmt.Println(unsafe.Alignof(v5)) // 8
}

// emptyStruct 空结构体示例
func emptyStruct() {
	// struct{}代表空结构体
	var v struct{}
	fmt.Println(unsafe.Sizeof(v)) // 0
}

// emptyStruct2 空结构体示例
func emptyStruct2() {
	// Test6自身的内存对齐要求是1字节
	// 空结构体（内存占用0）是最后一个字段时，
	// 编译器会额外再填充1字节
	type Test6 struct {
		a int8     // 1byte
		b struct{} // 0byte
	}

	var v6 Test6
	fmt.Println(unsafe.Sizeof(v6))  // 2
	fmt.Println(unsafe.Alignof(v6)) // 1

	fmt.Printf("v6:%p\n", &v6)     // v6:0xc00001a072
	fmt.Printf("v6.a:%p\n", &v6.a) // v6.a:0xc00001a072
	fmt.Printf("v6.b:%p\n", &v6.b) // v6.b:0xc00001a073

	// 空结构体（内存占用0）不是最后一个字段时，
	// 不会为空字段额外填充
	type Test7 struct {
		b struct{} // 0byte
		a int8     // 1byte
	}

	var v7 Test7
	fmt.Println(unsafe.Sizeof(v7))  // 1
	fmt.Println(unsafe.Alignof(v7)) // 1

	fmt.Printf("v7:%p\n", &v7)     // v7:0xc00001a090
	fmt.Printf("v7.a:%p\n", &v7.a) // v7.a:0xc00001a090
	fmt.Printf("v7.b:%p\n", &v7.b) // v7.b:0xc00001a090
}

// structPointer1 创建结构体指针示例
func structPointer1() {
	s1 := &Student{
		ID:     1,
		Name:   "七米",
		Gender: "男",
		Age:    18,
	}
	fmt.Println(s1) // &{1 七米 男 18}
}

// structPointer2 new函数创建结构体指针示例
func structPointer2() {
	var s1 = new(Student)

	(*s1).ID = 1
	(*s1).Name = "七米"
	(*s1).Gender = "男"
	(*s1).Age = 18

	s1.ID = 1
	s1.Name = "七米"
	s1.Gender = "男"
	s1.Age = 18
	fmt.Println(s1) // &{1 七米 男 18}
}

// NewStudent Student构造函数
func NewStudent(id, age int, name, gender string) *Student {
	return &Student{
		ID:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
	}
}

func createStudent() {
	s2 := NewStudent(1, 18, "张三", "19")
	fmt.Println(s2) // &{1 张三 19 18}
}

// structNest1 结构体嵌套示例
func structNest1() {

	// Teacher 老师
	type Teacher struct {
		Name   string
		Gender string
		Info   Info
	}

	t1 := Teacher{
		Name:   "Pony",
		Gender: "男",
		Info: Info{
			Email: "pony@test.com",
			Phone: "123456",
		},
	}
	fmt.Println(t1.Name)       // Pony
	fmt.Println(t1.Info.Email) // pony@test.com
}

// structNest2 结构体嵌套示例
func structNest2() {
	// Teacher 老师
	type Teacher struct {
		Name   string
		Gender string
		Info
	}

	var t2 Teacher

	t2.Name = "Bob"
	t2.Gender = "男"
	t2.Email = "bob@test.com" // 等价于t2.Info.Email="bob@test.com"
	t2.Phone = "654321"       // 等价于t2.Info.Phone="654321"

	t3 := Teacher{
		Name: "Tom",
		Info: Info{ // 字段名与结构体名称相同
			Email: "tom@test.com",
		},
	}
	fmt.Println(t3)

}

// structNest3 结构体嵌套示例
func structNest3() {

	type x struct {
		a int
	}

	type y struct {
		a int
		x // 嵌套x
	}

	var v y
	v.a = 100
	v.x.a = 200 // x不能省略

	fmt.Println(v)
	fmt.Println(v)
}

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

	for _, stu := range s {
		m[stu.name] = &stu
	}

	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}

func main() {
	//structDemo1()
	//structDemo2()
	//structDemo3()
	//structDemo4()
	//structDemo5()
	//structDemo6()
	//structDemo7()
	//structDemo8()
	//structDemo9()
	//structDemo10()

	//emptyStruct2()

	//structPointer1()
	//structPointer2()

	//createStudent()

	//structNest1()
	//structNest2()
	structNest3()

	//structEx1()

}
