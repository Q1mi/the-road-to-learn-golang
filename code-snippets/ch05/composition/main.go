package main

import (
	"fmt"
	"sync"
)

// Info 联系方式信息
type Info struct {
	Email string
	Phone string
}

// Detail 详细联系信息
func (i Info) Detail() string {
	return fmt.Sprintf("邮箱(%s) 电话(%s)",
		i.Email, i.Phone,
	)
}

// Teacher 老师
type Teacher struct {
	Name   string
	Gender string
	Info
}

// Introduce 进行个人介绍的方法
func (t Teacher) Introduce() {
	fmt.Println("姓名：", t.Name)
	fmt.Println("性别：", t.Gender)
	fmt.Println("联系方式：", t.Info.Detail())
}

//var (
//	mux     sync.Mutex                // 防止数据竞争的互斥锁
//	mapping = make(map[string]string) // 缓存的存储
//)
//
//// Query 根据指定key查询缓存
//func Query(key string) string {
//	mux.Lock() // 加锁
//	v := mapping[key]
//	mux.Unlock() // 释放锁
//	return v
//}

// cache 定义一个表示缓存的匿名结构体变量
var cache = struct {
	sync.Mutex // 嵌入
	mapping    map[string]string
}{
	mapping: make(map[string]string),
}

// Query 根据指定key查询缓存
func Query(key string) string {
	cache.Lock() // 加锁
	v := cache.mapping[key]
	cache.Unlock() // 解锁
	return v
}

func main() {
	t := Teacher{
		Name:   "Tom",
		Gender: "男",
	}

	t.Email = "tom@test.com"
	fmt.Println(t.Info.Email) // tom@test.com
	t.Info.Phone = "13245768"
	fmt.Println(t.Phone) // 13245768

	fmt.Println(t.Info.Detail()) // 邮箱(tom@test.com) 电话(13245768)
	fmt.Println(t.Detail())      // 邮箱(tom@test.com) 电话(13245768)

	t.Introduce()
}
