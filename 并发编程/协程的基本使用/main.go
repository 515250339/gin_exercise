package main

import (
	"fmt"
	"time"
)

func main() {
	go test() //表示开启一个协程
	for i := 0; i < 2; i++ {
		fmt.Println("main() 你好")
		time.Sleep(time.Millisecond * 100)
	}
}

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test() 你好")
		time.Sleep(time.Millisecond * 100)
	}
}
