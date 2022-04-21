package main

import "fmt"

func test1(ch chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println("test1")
		ch <- 1
	}
	close(ch)
}

func test2(ch1 chan int, ch2 chan int) {
	for {
		i, ok := <-ch1 // 通道关闭后再取值ok=false
		fmt.Println("test2")
		if !ok {
			fmt.Println("test2--break")
			break
		}
		ch2 <- i * i
	}
	close(ch2)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 开启goroutine将0~10的数发送到ch1中
	go test1(ch1)
	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go test2(ch1, ch2)
	// 在主goroutine中从ch2中接收值打印
	for i := range ch2 { // 通道关闭后会退出for range循环
		fmt.Println(i)
	}
}
