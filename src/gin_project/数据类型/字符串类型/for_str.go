package main

import "fmt"

func main() {
	s := "hello word"
	// 通过下标去获取元素
	for i := 0; i < len(s); i++ {
		//%v 相应值的默认格式; %c 输出单个字符
		fmt.Printf("%v(%c)   ", s[i], s[i])
	}
	fmt.Println()
	// 遍历循环元素
	for _, j := range s {
		fmt.Printf("%v(%c)   ", j, j)
	}

}
