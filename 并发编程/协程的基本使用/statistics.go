package main

import (
	"fmt"
	"sync"
	"time"
)

var wg2 sync.WaitGroup

func test3(n int) {
	for num := (n-1)*30000 + 1; num < n*30000; num++ {
		if num > 1 {
			var flag = true
			for i := 2; i < num; i++ {
				if num%i == 0 {
					flag = false
					break
				}
			}
			if flag {
				fmt.Println(num, "是素数")
			}
		}
	}
	wg2.Done()
}

func main() {
	start := time.Now().Unix()
	for i := 1; i <= 4; i++ {
		wg2.Add(1)
		go test3(i)
	}
	wg2.Wait()
	fmt.Println("执行完毕")
	end := time.Now().Unix()
	fmt.Println(end - start)
}
