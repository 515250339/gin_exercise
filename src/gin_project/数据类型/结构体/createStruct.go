package main

import "fmt"

type person struct {
	name string
	age  int
	sex  string
}

func main() {
	var p1 person
	p1.name = "张三"
	p1.age = 15
	p1.sex = "男"
	fmt.Printf("type:%T  value:%[1]v \n", p1)
	fmt.Printf("type:%T  value:%#[1]v \n", p1)
}
