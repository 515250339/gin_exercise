package main

import (
	"errors"
	"fmt"
)

func test1() {
	panic("panic in 1")
}

func test2() {
	test1()
}

func funcA() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println(res)
}

func readFile(fileName string) error {
	if fileName == "main.go" {
		return nil
	}
	return errors.New("读取文件错误")
}

func main() {
	//test2()

	//funcA()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	var err = readFile("main.go")
	if err != nil {
		println(err)
	}
	fmt.Println("继续执行")

}
