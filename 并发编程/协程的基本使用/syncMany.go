package main

import (
	"fmt"
	"sync"
)

var wg1 sync.WaitGroup

func hello(i int) {
	defer wg1.Done() // goroutine结束就登记-1
	fmt.Println("hello world", i)
}

func main() {
	for i := 0; i < 10; i++ {
		wg1.Add(1) // 启动一个goroutine就登记+1
		go hello(i)
	}
	wg1.Wait() // 等待所有登记的goroutine都结束
}
