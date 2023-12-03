package main

import (
	"fmt"
	"os"
	"time"
)

func readFromFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close() // 函数结束时关闭文件

	// 从文件中读取内容的操作...
	return nil
}

// recordTime 打印函数执行时间
func recordTime(funcName string) func() {
	start := time.Now()
	fmt.Printf("enter %s\n", funcName)
	return func() {
		fmt.Printf("exit %s, cast:%v\n", funcName, time.Since(start))
	}
}

func slowFunc() {
	defer recordTime("slowFunc")()
	time.Sleep(time.Second * 10)
}
