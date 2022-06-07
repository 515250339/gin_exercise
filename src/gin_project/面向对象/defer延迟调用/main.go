package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	//fmt.Println("start")
	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//defer fmt.Println(3)
	//fmt.Println("end")

	//var x, y int
	//defer calc("AA", x, calc("A", x, y))
	//x = 10
	//defer calc("BB", x, calc("B", x, y))
	//y = 10

	var whatever [5]struct{}
	for i := range whatever {
		defer func() {
			fmt.Println(i)
		}()
	}
}
