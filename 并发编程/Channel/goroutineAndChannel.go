package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	intCh := make(chan int, 10)
	wg.Add(2)
	go write(intCh)
	go read(intCh)
	wg.Wait()
	fmt.Println("读取完毕")
}

func write(intCh chan int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		intCh <- i
	}
	close(intCh)
}

func read(intCh chan int) {
	defer wg.Done()
	for v := range intCh {
		fmt.Println(v)
		time.Sleep(time.Second)
	}
}
