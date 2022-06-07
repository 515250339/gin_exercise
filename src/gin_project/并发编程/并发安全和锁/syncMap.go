package main

import (
	"fmt"
	"strconv"
	"sync"
)

var m = make(map[string]int)
var wg4 sync.WaitGroup

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}
func main() {
	//for i := 0; i < 10; i++ {
	//	wg4.Add(1)
	//	go func(n int) {
	//		key := strconv.Itoa(n)
	//		set(key, n)
	//		fmt.Printf("k = %v, v = %v \n", key, get(key))
	//		wg4.Done()
	//	}(i)
	//}
	//wg4.Wait()

	var m2 = sync.Map{}
	for i := 0; i < 10; i++ {
		wg4.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m2.Store(key, n)
			value, _ := m2.Load(key)
			fmt.Printf("k = %v, v = %v \n", key, value)
			wg4.Done()
		}(i)
	}
	wg4.Wait()
}
