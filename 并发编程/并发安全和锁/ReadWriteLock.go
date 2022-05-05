package main

import (
	"fmt"
	"sync"
	"time"
)

var wg2 sync.WaitGroup
var mutex sync.Mutex

func write() {
	mutex.Lock()
	fmt.Println("执行写的操作")
	time.Sleep(time.Second * 1)
	mutex.Unlock()
	wg2.Done()
}

func read() {
	mutex.Lock()
	fmt.Println("执行读的操作")
	time.Sleep(time.Second * 1)
	mutex.Unlock()
	wg2.Done()
}

func main() {
	for i := 0; i < 10; i++ {
		wg2.Add(1)
		go write()
	}
	for i := 0; i < 10; i++ {
		wg2.Add(1)
		go read()
	}
	wg2.Wait()
}
