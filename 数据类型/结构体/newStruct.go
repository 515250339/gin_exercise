package main

import "fmt"

type person2 struct {
	name string
	age  int
	sex  string
}

type person3 struct {
	name string
	age  int
	string
}

func main() {
	//var p1 = new(person2)
	//p1.name = "zs"
	//p1.age = 111
	//p1.sex = "女"
	//fmt.Printf("type:%T value:%[1]v \n", p1)
	//fmt.Printf("type:%T value:%#[1]v \n", p1)

	//var p1 = &person2{
	//	name: "张三",
	//	age:  55,
	//	sex:  "中",
	//}
	//(*p1).age = 40 //这样也是可以的
	//fmt.Printf("type:%T value:%[1]v \n", p1)
	//fmt.Printf("type:%T value:%#[1]v \n", p1)
	//p1 := person2{
	//	name: "张三",
	//	age:  55,
	//	sex:  "中",
	//}
	//fmt.Printf("type:%T value:%[1]v \n", p1)
	//fmt.Printf("type:%T value:%#[1]v \n", p1)

	//p1 := &person2{
	//	name: "张三",
	//	age:  55,
	//	sex:  "中",
	//}
	//(*p1).name = "qqq"
	//fmt.Printf("type:%T value:%[1]v \n", *p1)
	//fmt.Printf("type:%T value:%#[1]v \n", *p1)

	p1 := person3{
		"张三",
		55,
		"嬲",
	}
	fmt.Printf("type:%T value:%[1]v \n", p1)
	fmt.Printf("type:%T value:%#[1]v \n", p1)
}
