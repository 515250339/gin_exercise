package main

import (
	"bufio"
	"fmt"
	"os"
)

func test1() {
	// 定义一个 map 对象 key:str value:int
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "end" {
			break
		}
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func test2() {
	// 创建 map
	counts := make(map[string]int)
	// 切片截取命令行输入数据
	files := os.Args[1:]
	// 判断命令行是否输入文件名
	if len(files) == 0 {
		// 无输入则调用 countLines 方法
		countLines(os.Stdin, counts)
	} else {
		// 输入则打开文件 支持多个文件
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Println(os.Stderr, "dup2: %v\n", err)
			}
			// 判断有无重复数据
			countLines(f, counts)
			// 关闭文件
			err = f.Close()
			if err != nil {
				return
			}
		}
	}
	for line, n := range counts {
		// 打印 counst数据
		fmt.Printf("%d\t%s\n", n, line)
	}

}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func main() {
	test2()
}
