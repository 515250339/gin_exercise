package main

import "fmt"

type Person4 struct {
	name string
	age  int
}

type room3 struct {
	name string
}

type floor3 struct {
	name  string
	room3 // 通过嵌套匿名结构体实现继承
}

type floor4 struct {
	name   string
	*room3 //通过嵌套匿名结构体实现继承
}

type myInt int

func (i *myInt) myType() {
	fmt.Printf("type:%T value:%#[1]v \n", *i)
}

//值类型接受者
func (p Person4) personName() {
	p.name = "张三"
	fmt.Printf("type:%T value:%#[1]v \n", p)

}

// 指针类型接受者
func (p *Person4) personAge() {
	p.age = 998
	fmt.Printf("type:%T value:%#[1]v \n", p)

}

func (r *room3) room() {
	r.name = "301"
}

func (f *floor3) floor() {
	f.name = "3层"
}

func (f *floor4) floor4() {
	f.name = "4层"
}

func main() {
	//p := Person4{}
	//p.personAge()
	//p.personName()

	//f := floor3{}
	//f.room()
	//f.floor()
	//fmt.Printf("%v的%v室", f.name, f.room3.name)

	//f := &floor4{
	//	name: "",
	//	room3: &room3{
	//		"403",
	//	},
	//}
	//f.floor4()
	//fmt.Printf("%v的%v室", f.name, f.room3.name)

	var i myInt
	i = 100
	i.myType()
}
