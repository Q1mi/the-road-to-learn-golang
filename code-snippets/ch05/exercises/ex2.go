package main

import (
	"fmt"
	"io"
)

// 编写一个学生信息管理系统。
// 学生有id、姓名、年龄、分数等信息
// 程序提供展示学生列表、添加学生、编辑学生信息、删除学生等功能

// student 学生
type student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// studentMgr 学生管理
type studentMgr struct {
	students map[int]student
}

// showStudents 展示所有学生
func (sm *studentMgr) showStudents() {
	fmt.Println("所有的学员如下：")
	for _, v := range sm.students {
		fmt.Println(v)
	}
}

// addStudent 添加学生
func (sm *studentMgr) addStudent(stu student) {
	_, ok := sm.students[stu.ID]
	if ok {
		fmt.Println("该学生已存在！")
	} else {
		sm.students[stu.ID] = stu
		fmt.Println("添加学生成功！")
	}
}

// delStudent 根据id删除学生
func (sm *studentMgr) delStudent(id int) {
	_, ok := sm.students[id]
	if ok {
		delete(sm.students, id)
		fmt.Println("删除学生成功！")
	} else {
		fmt.Println("查无此人！")
	}
}

// editStudent 编辑学生
func (sm *studentMgr) editStudent(id int, name string) {
	_, ok := sm.students[id]
	if ok {
		stu := student{
			id,
			name,
		}
		sm.students[id] = stu
		fmt.Printf("学号: %d 名字修改为：%s", id, name)
	} else {
		fmt.Println("查无此人！")
	}
}

func main() {
	sm := studentMgr{
		students: make(map[int]student, 8),
	}
loop:
	for {
		fmt.Println(`
		 1.添加
		 2.删除
		 3.修改
		 4.显示
		 5.退出
		`)
		fmt.Print("请输入要执行的菜单序号：")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			var id int
			var name string
			fmt.Print("请输入学号：")

			fmt.Scanln(&id)
			fmt.Print("请输入姓名：")
			fmt.Scanln(&name)
			std := student{
				ID:   id,
				Name: name,
			}
			sm.addStudent(std)
		case 2:
			fmt.Print("请输入要删除的学号：")
			var id int
			fmt.Scanln(&id)
			sm.delStudent(id)

		case 3:
			var (
				id   int
				name string
			)
			fmt.Print("请输入要修改的学号：")
			fmt.Scanln(&id)
			fmt.Print("名字修改为：")
			fmt.Scanln(&name)
			sm.editStudent(id, name)
		case 4:
			sm.showStudents()
		case 5:
			break loop
		default:
			fmt.Println("请输入正确的选项！")
		}
	}
	io.Copy()
}
