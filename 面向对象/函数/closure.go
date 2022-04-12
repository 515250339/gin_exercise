package main

import (
	"fmt"
	"strings"
)

func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func adder2(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}

}

func makeSuffixFunc(suffix string) func(string2 string) string {

	return func(name string) string {
		// 判断 name 是否以 suffix 结尾
		// 判断 name 是否以 suffix 结尾
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}

}

func calcTest(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub

}

func main() {
	//var f = adder()
	//fmt.Println(f(10))
	//fmt.Println(f(20))
	//fmt.Println(f(30))
	//
	//f1 := adder()
	//fmt.Println(f1(30))
	//fmt.Println(f1(30))

	//var f = adder2(1)
	//fmt.Println(f(10))
	//fmt.Println(f(20))
	//fmt.Println(f(30))
	//
	//f1 := adder2(2)
	//fmt.Println(f1(30))
	//fmt.Println(f1(30))

	//jpg := makeSuffixFunc(".jpg")
	//txt := makeSuffixFunc(".txt")
	//fmt.Println(jpg("test"))
	//fmt.Println(txt("test.txt"))

	f1, f2 := calcTest(2)
	fmt.Println(f1(1), f2(1))
}
