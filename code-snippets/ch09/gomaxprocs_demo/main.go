package main

import (
	"fmt"
)

func a() {
	for i := 1; i < 20; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 20; i++ {
		fmt.Println("B:", i)
	}
}

//func main() {
//	runtime.GOMAXPROCS(1)
//	wg := sync.WaitGroup{}
//	wg.Add(20)
//	for i := 0; i < 10; i++ {
//		go func(i int) {
//			fmt.Println("A: ", i)
//			wg.Done()
//		}(i)
//	}
//	for i := 0; i < 10; i++ {
//		go func(i int) {
//			fmt.Println("B: ", i)
//			wg.Done()
//		}(i)
//	}
//	wg.Wait()
//}

func main() {
	go fmt.Println("hi")
	for {
	}
}
