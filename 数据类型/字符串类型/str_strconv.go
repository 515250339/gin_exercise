package main

import (
	"fmt"
	"strconv"
)

func main() {
	// int 转换为 string
	var i int
	fmt.Printf("转换前类型：%T，值：%v \n", i, i)
	s1 := strconv.Itoa(i)
	fmt.Printf("转换后类型：%T，值：%v \n", s1, s1)
	fmt.Println()

	// float 转 string
	var i2 float64 = 0.01
	fmt.Printf("转换前类型：%T，值：%v \n", i2, i2)
	/* 参数 1：要转换的值
	参数 2：格式化类型
	参数 3: 保留的小数点 -1（不对小数点格式化）
	参数 4：格式化的类型
	*/
	s2 := strconv.FormatFloat(i2, 'f', 2, 64)
	fmt.Printf("转换后类型：%T，值：%v \n", s2, s2)
	fmt.Println()

	// bool 转 string
	b := false
	fmt.Printf("转换前类型：%T，值：%v \n", b, b)
	s3 := strconv.FormatBool(b)
	fmt.Printf("转换后类型：%T，值：%v \n", s3, s3)
	fmt.Println()

	// int64 转 string
	var i3 int64 = 998
	fmt.Printf("转换前类型：%T，值：%v \n", i3, i3)
	var s4 = strconv.FormatInt(i3, 10)
	fmt.Printf("转换后类型：%T，值：%v \n", s4, s4)
	fmt.Println()

}
