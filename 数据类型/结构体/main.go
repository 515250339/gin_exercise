package main

import "fmt"

type name string // 类型定义
type age = int   // 类型别名

type room struct {
	name string
}

type floor struct {
	name string
	room room
}

type building struct {
	name  string
	floor floor
}

type Room2 struct {
	name string
}

type Floor2 struct {
	name string
	Room2
}

type Building2 struct {
	name string
	Floor2
}

func structTest1() {
	info := building{
		"大厦A",
		floor{
			"2楼",
			room{
				"A-2-3",
			},
		},
	}
	fmt.Printf("type:%T value:%[1]v \n", info)
	fmt.Printf("type:%T value:%#[1]v \n", info)

}

func structTest2() {
	// 需先初始化
	//info := new(Building2)
	info := Building2{}
	info.name = "B"
	info.Floor2.name = "5层"
	info.Room2.name = "B-5-6"
	fmt.Printf("type:%T value:%[1]v \n", info)
	fmt.Printf("type:%T value:%#[1]v \n", info)

}

func (r room) structFunc() {
	//p1 := p{
	//	name: "",
	//	age:  0,
	//	sex:  "",
	//}
	fmt.Printf("type:%T value:%[1]v \n", r)
	fmt.Printf("type:%T value:%#[1]v \n", r)
}

func main() {
	//var a_name name
	//a_name = "张三"
	//var a_age age
	//a_age = 11
	//fmt.Printf("type: %T  value:%[1]s \n", a_name)
	//fmt.Printf("type: %T  value:%[1]v \n", a_age)
	//structTest1()
	//structTest2()
	r := room{
		name: "小王子",
	}
	r.structFunc()
}
