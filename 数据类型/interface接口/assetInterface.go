package main

import "fmt"

func main() {
	var x interface{}
	x = 123
	v, ok := x.(string)
	if ok {
		fmt.Println("ok", v, ok)
	} else {
		fmt.Println("no", v, ok)
	}
}
