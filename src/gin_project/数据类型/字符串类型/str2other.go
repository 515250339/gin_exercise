package main

import (
	"fmt"
	"strconv"
)

// string转int
func str2int2() {
	var a string = "1234"
	i64, _ := strconv.ParseInt(a, 10, 64)
	fmt.Printf("转换前类型: %T 转换后类型: %T 值: %v \n", a, i64, i64)
}

// string转float
func str2float() {
	var a string = "3.14"
	f32, _ := strconv.ParseFloat(a, 32)
	f64, _ := strconv.ParseFloat(a, 64)
	fmt.Printf("转换前类型: %T 转换后类型: %T 值: %v \n", a, f32, f32)
	fmt.Printf("转换前类型: %T 转换后类型: %T 值: %v \n", a, f64, f64)
}

// string转bool
func str2bool() {
	var a string = "true"
	b, _ := strconv.ParseBool(a)
	fmt.Printf("转换前类型: %T 转换后类型: %T 值: %v \n", a, b, b)
}

// string转字符
func str2ASCII() {
	var a string = "hello goland"
	for _, i := range a {
		fmt.Printf("%v(%c)  ", i, i)
	}
}

// 字符串反转
func strServer() {
	a := "test"
	r := []rune(a)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	fmt.Printf(string(r))
}

func main() {
	//str2int2()
	//str2float()
	//str2bool()
	//str2ASCII()
	strServer()
}
