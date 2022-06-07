package main

import "fmt"

func inSum(a, b int) int {
	return a + b
}

func inSum2(a string, x ...int) int {
	fmt.Println(a) // zs
	fmt.Println(x) // x是一个切片 = [1, 2]
	sum := 0
	for _, i := range x {
		sum += i
	}
	return sum
}

func calc(a, b int) (int, int) {
	var sum = a + b
	var sub = a - b
	return sum, sub

}

type test func(int, int) int

const age int = 998

func test2() {
	name := "张三"
	fmt.Println(name)
}
func test3(x, y int) {
	fmt.Println(x, y) //函数的参数也是只在本函数中生效
	if x < y {
		z := 9 // 变量 z 只在 if 语句块生效
		fmt.Println(z)
	}
	//fmt.Println(z)
}

func forTest() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	//fmt.Println(i)
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func do(s string) func(int, int) int {
	switch s {
	case "+":
		return add
	case "-":
		return sub
	default:
		return nil
	}
}

func main() {
	//ret := inSum(1, 2)
	//fmt.Println(ret)

	//str := "zs"
	//ret := inSum2(str, 1, 2)
	//fmt.Println(ret)

	//sum, sub := calc(1, 2)
	//fmt.Println(sum, sub)

	//var t test // 声明一个 test 类型的变量 t
	//t = inSum  // 把 inSum 赋值给 t
	//fmt.Println(t(1, 3))
	//fmt.Printf("type of t:%T\n", t)

	//fmt.Printf("type of age:%T  value: %[1]v \n", age)

	//fmt.Printf("type of age:%T  value: %[1]v \n", name)

	//test3(3, 8)
	//forTest()

	//a := do("+")
	//b := do("-")
	//fmt.Println(a(1, 3))
	//fmt.Println(b(1, 3))

	// 一：匿名函数  匿名自执行函数
	func() {
		fmt.Println("test")
	}()
	// 二：匿名函数
	var fn = func(a, b int) int {
		return a + b
	}
	fmt.Println(fn(1, 2))
	// 三：匿名函数 自动执行 接收参数
	func(x, y int) {
		fmt.Println(x, y)
	}(10, 20)

}
