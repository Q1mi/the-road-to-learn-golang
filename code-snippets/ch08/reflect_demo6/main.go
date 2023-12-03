package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// 结构体反射练习

// Student 学生结构体
type Student struct {
	Name string `info:"name"`
	Age  int    `info:"age"`
}

// 给Student结构体添加Study和Sleep两个方法
// 注意：方法名称首字母大写

// Study 学习
func (s Student) Study(title string) {
	fmt.Printf("同学%s正在学习%s\n", s.Name, title)
}

// Play 玩
func (s Student) Play(hours int) {
	fmt.Printf("同学%s玩了%d小时\n", s.Name, hours)
}

// Do 调用变量v的name方法
func Do(v interface{}, name string, arg interface{}) {
	tInfo := reflect.TypeOf(v)
	vInfo := reflect.ValueOf(v)

	fmt.Println(tInfo.NumMethod())
	m := vInfo.MethodByName(name)
	//m, has := tInfo.MethodByName(name)
	if !m.IsValid() || m.IsNil() {
		fmt.Printf("%s没有%s方法\n", tInfo.Name(), name)
		return
	}

	// 调用指定方法
	// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
	argVal := reflect.ValueOf(arg)
	m.Call([]reflect.Value{argVal})
}

// LoadInfo 加载数据至变量v
func LoadInfo(s string, v interface{}) (err error) {
	// 确保传入的v是结构体指针
	tInfo := reflect.TypeOf(v)
	if tInfo.Kind() != reflect.Ptr {
		err = errors.New("please pass into a struct ptr")
		return
	}
	if tInfo.Elem().Kind() != reflect.Struct {
		err = errors.New("please pass into a struct ptr")
		return
	}

	vInfo := reflect.ValueOf(v)

	// 按行分隔
	list := strings.Split(s, "\n")
	for _, item := range list {
		// 按等号拆分为key和value的配置对
		kvList := strings.Split(item, "=")
		if len(kvList) != 2 {
			continue
		}
		fieldName := ""
		key := strings.TrimSpace(kvList[0])
		value := strings.TrimSpace(kvList[1])
		// 遍历结构体字段的tag找到对应的key
		for i := 0; i < tInfo.Elem().NumField(); i++ {
			f := tInfo.Elem().Field(i)
			tagVal := f.Tag.Get("info")
			if tagVal == key {
				fieldName = f.Name
				break
			}
		}
		if len(fieldName) == 0 {
			continue // 找不到跳过
		}

		// 根据找到的结构体字段名称找到结构体的字段
		fv := vInfo.Elem().FieldByName(fieldName)
		switch fv.Type().Kind() {
		case reflect.String:
			fv.SetString(value)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			intVal, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return err
			}
			fv.SetInt(intVal)
		default:
			return fmt.Errorf("unsupport value type:%v", fv.Type().Kind())
		}
	}
	return
}

func main() {
	var stu Student
	// 从文本文件中读取内容
	s, err := os.ReadFile("info.txt")
	if err != nil {
		panic(err)
	}

	err = LoadInfo(string(s), &stu)
	fmt.Printf("stu:%#v err:%v\n", stu, err)

	Do(stu, "Study", "《跟七米学Go语言》")
}
