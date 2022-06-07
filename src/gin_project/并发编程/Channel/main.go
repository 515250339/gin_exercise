package main

import "fmt"

func main() {
	//ch := make(chan int)
	//ch <- 10
	//fmt.Println("发送成功")

	ch := make(chan int, 5)
	ch <- 10
	ch <- 12
	fmt.Println("发送成功")
}
