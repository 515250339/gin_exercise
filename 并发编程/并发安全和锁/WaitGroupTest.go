package main

import (
	"fmt"
	"sync"
)

var wg3 sync.WaitGroup

func hello() {
	defer wg3.Done()
	fmt.Println("Hello Goroutine!")
}

func main() {
	wg3.Add(1)
	go hello() // 启动另外一个goroutine去执行hello函数
	fmt.Println("main goroutine done !")
	wg3.Wait()
}
