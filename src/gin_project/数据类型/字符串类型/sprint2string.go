package main

import "fmt"

func main() {
	var i int = 10
	var f float64 = 2.00
	var t bool = true
	var b byte = 'a'
	var str string

	str = fmt.Sprintf("%d", i)
	fmt.Printf("类型: %T，值: %v \n", str, str)

	str = fmt.Sprintf("%f", f)
	fmt.Printf("类型: %T，值: %v \n", str, str)

	str = fmt.Sprintf("%t", t)
	fmt.Printf("类型: %T，值: %v \n", str, str)

	str = fmt.Sprintf("%c", b)
	fmt.Printf("类型: %T，值: %v \n", str, str)
}
