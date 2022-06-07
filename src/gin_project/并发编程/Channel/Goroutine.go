package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	// 开启三个 goroutine

	for w := 0; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 5个任务
	for w := 0; w <= 5; w++ {
		jobs <- w
	}

	close(jobs)

	// 输出结果
	for a := 1; a <= 5; a++ {
		<-results
	}
}
