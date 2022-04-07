package main

import "fmt"

type Append interface {
	add()
}

type Delete interface {
	pop()
}

// SliceMethod 接口嵌套
type SliceMethod interface {
	Append
	Delete
}

// Slice 创建切片
type Slice []int

// 给切片添加元素
func (s *Slice) add() {
	*s = append(*s, 6)
}

// 给切片删除元素
func (s *Slice) pop() {
	*s = append((*s)[:2], (*s)[4:]...)
}

func run() *Slice {
	// 初始化数据
	s := &Slice{1, 2, 3, 4, 5}
	fmt.Println(*s)
	// 接口赋值
	var sm SliceMethod = s
	// 执行添加/删除方法
	sm.add()
	sm.pop()
	return s
}

func main() {
	s := run()
	fmt.Println(*s)
}
