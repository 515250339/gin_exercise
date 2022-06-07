package main

import (
	"fmt"
	"time"
)

func timer(fn func()) func() {
	return func() {
		startTime := time.Now().Unix()
		fn()
		endTime := time.Now().Unix()
		fmt.Println("运行时间: ", endTime-startTime)
	}
}

func testFunc() {
	fmt.Println("运行 testFunc")
	time.Sleep(time.Second * 2)
}

func main() {
	test := timer(testFunc)
	test()
}
