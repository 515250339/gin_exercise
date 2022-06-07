package main

import "fmt"

func pointer() {
	a := 10
	// 值
	fmt.Printf("%d \n", a)
	// 指针地址
	fmt.Printf("%d \n", &a)
	// 去内存地址取值
	fmt.Printf("%d \n", *&a)
	// 指针类型
	fmt.Printf("%T \n", &a)
}

func pointer2() {
	a := 10
	b := &a
	fmt.Printf("%d prt: %p \n", a, &a)
	fmt.Printf("%v type: %T \n", b, b)
	fmt.Printf("取b的地址 %d", &b)
}

func pointer3() {
	a := 10
	b := &a
	fmt.Printf("type: %T \n", b)
	c := *b
	fmt.Printf("value: %d type: %[1]T \n", c)
}
func pointer4(x int) {
	x = 9
	fmt.Printf("value: %d type: %[1]T \n", x)
}

func pointer5(x *int) {
	*x = 9
	fmt.Printf("value: %d type: %[1]T \n", x)
	fmt.Printf("value: %d type: %[1]T \n", *x)
}

func createMake() {
	a := make(map[string]string)
	a["name"] = "张三"
	fmt.Println(a)

	b := make([]int, 2, 5)
	b = append(b, 5)
	fmt.Println(b)

	b1 := new([]int)
	*b1 = append(*b1, 3)
	fmt.Println(b1)

}

func createNew() {
	// 实例化int
	a := new(int)
	*a = 10
	fmt.Printf("value %d, type %[1]T \n", *a)
	// 实例化slice
	b := new([]int)
	*b = append(*b, 3)
	fmt.Printf("value %d, type %[1]T \n", *b)
	// 实例化map
	c := new(map[string]string)
	*c = map[string]string{}
	(*c)["name"] = "张三"
	fmt.Printf("value %v, type %[1]T \n", *c)
}

type people struct {
	name string
	age  int
}

func main() {
	//var a = 10
	//pointer4(a)
	//fmt.Println(a)
	//pointer5(&a)
	//fmt.Println(a)

	//var a map[string]string
	//a["name"] = "张三"
	//fmt.Println(a)

	//var a *people
	//// 分配空间
	//a = new(people)
	//a.name = "张三"
	//fmt.Printf("value %v, type %[1]T \n", a)

	a := make([]int, 2, 4)
	b := make(map[string]int)
	var c = make(chan int, 1)
	fmt.Printf("value %v, type %[1]T \n", a)
	fmt.Printf("value %v, type %[1]T \n", b)
	fmt.Printf("value %v, type %[1]T \n", c)
}
