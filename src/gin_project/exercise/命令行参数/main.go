package main

import (
	"fmt"
	"os"
	"strings"
)

/*
练习 1.1 修改 echo 程序输出 os.Args[0] ,即命令的名字
练习 1.2 修改 echo 输出参数的索引和值，每行一个
练习1.3：尝试测量可能低效的程序和使用 strings.Join 的程序在执行时间上的差昇
（1.6节有time 包，11.4 节展示如何撰写系统性的性能评估测试。
*/

// 练习 1.1
func test1() {
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = "****"
	}
	fmt.Println(s)
}

//练习 1.2
func test2() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = "****"
		fmt.Println(i, s)
	}
}

//练习1.3
func test3() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(strings.Join(os.Args[i:], "+++"))
	}

}
func main() {
	test1()
	fmt.Println()
	test2()
	fmt.Println()
	test3()
}
