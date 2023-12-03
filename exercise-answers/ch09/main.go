package main

import (
	"fmt"
	"math/rand"
)

/*
1. 使用 goroutine 和 channel 实现一个计算int64随机数各位数和的程序，例如，生成随机数61345，计算其每位上的数字之和为19。
2. 开启一个 goroutine 循环生成int64类型的随机数，发送到jobChan。
3. 开启24个 goroutine，从jobChan中取出随机数计算各位数的和，将结果发送到resultChan。
4. 主 goroutine 从resultChan取出结果并在终端输出。
*/

type Job struct {
	Number int
	Id     int
}

type Result struct {
	job *Job
	sum int
}

func calc(number int) (sum int) {
	for number != 0 {
		tmp := number % 10
		sum += tmp
		number /= 10
	}
	return
}

// consumer 消费者
func consumer(jobChan <-chan *Job, resultChan chan<- *Result) {
	for job := range jobChan {
		sum := calc(job.Number)
		r := &Result{
			job: job,
			sum: sum,
		}
		resultChan <- r
	}
}

func startWorkerPool(num int, jobChan <-chan *Job, resultChan chan<- *Result) {
	for i := 0; i < num; i++ {
		go consumer(jobChan, resultChan)
	}
}

// producer 生产者
func producer(jobChan chan<- *Job) {

	id := 1
	for {
		id++
		number := rand.Int()
		job := &Job{
			Id:     id,
			Number: number,
		}
		jobChan <- job
	}
}

func printResult(resultChan <-chan *Result) {
	for result := range resultChan {
		fmt.Printf("job id:%v number:%v result:%d\n", result.job.Id, result.job.Number, result.sum)
	}
}

func main() {
	jobChan := make(chan *Job, 1000)
	resultChan := make(chan *Result, 1000)

	go producer(jobChan)

	startWorkerPool(24, jobChan, resultChan)

	printResult(resultChan)
}
