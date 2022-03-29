package main

import (
	"fmt"
	"sort"
)

func createSlice() {
	// 切片是引用类型，不支持直接比较，只能和 nil 比较

	//声明一个字符串切片
	var a []string
	fmt.Println(a)
	fmt.Println(a == nil)

	//声明一个字符串切片并初始化
	var b = []string{}
	fmt.Println(b)
	fmt.Println(b == nil)

	//声明一个布尔切片并初始化赋值
	var c = []bool{false}
	fmt.Println(c)
	fmt.Println(c == nil)
}

// 切片的长度和容量
func showCap() {
	//声明一个字符串切片
	var a = []int{2, 3, 4, 5}
	fmt.Printf("长度%v 容量%v 值%v \n", len(a), cap(a), a)

	b := a[:2]
	fmt.Printf("长度%v 容量%v 值%v \n", len(b), cap(b), b)

	b = a[1:4]
	fmt.Printf("长度%v 容量%v 值%v \n", len(b), cap(b), b)
}

func createArraySlice() {
	var a = [5]int{1, 2, 3, 4, 5}
	b := a[1:3]
	fmt.Printf("类型%T value%[1]v \n", a)
	fmt.Printf("类型%T value%[1]v \n", b)
	c := b[:1]
	fmt.Printf("类型%T value%[1]v \n", c)
}

func makeCreateSlice() {
	// Type, size cap
	a := make([]int, 2, 10)
	fmt.Printf("长度%v 容量%v 值%v \n", len(a), cap(a), a)
}

func appendMethod() {
	a := make([]int, 2, 3)
	a = append(a, 10)
	fmt.Printf("长度%v 容量%v 值%v \n", len(a), cap(a), a)
	a = append(a, 100)
	fmt.Printf("长度%v 容量%v 值%v \n", len(a), cap(a), a)
}

func A() {
	a := make([]int, 2, 2)
	fmt.Printf("长度%v 容量%v 值%v \n", len(a), cap(a), a)
	b := []int{10, 11, 122, 123, 133}
	a = append(a, b...)
	// 如果最终容量（cap）计算值溢出，则最终容量（cap）就是新申请容量（cap）
	fmt.Printf("长度%v 容量%v 值%v \n", len(a), cap(a), a)
}

func delMethod() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("长度%v 容量%v 值%v \n", len(a), cap(a), a)
	a = append(a[:2], a[3:]...)
	fmt.Printf("长度%v 容量%v 值%v \n", len(a), cap(a), a)
}

func mergeSlice() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6, 7}
	c := append(a, b...)
	fmt.Printf("长度%v 容量%v 值%v \n", len(c), cap(c), c)
}

func copySlice() {
	a := []int{1, 2, 3, 4}
	fmt.Printf("长度%v 容量%v 值%v \n", len(a), cap(a), a)
	b := a
	fmt.Printf("长度%v 容量%v 值%v \n", len(b), cap(b), b)
	b[0] = 998
	fmt.Printf("a 长度%v 容量%v 值%v \n", len(a), cap(a), a)
	fmt.Printf("b 长度%v 容量%v 值%v \n", len(b), cap(b), b)

}

func deepcopySlice() {
	a := []int{1, 2, 3, 4}
	fmt.Printf("长度%v 容量%v 值%v \n", len(a), cap(a), a)
	b := []int{5, 6, 7, 8}
	fmt.Printf("长度%v 容量%v 值%v \n", len(b), cap(b), b)
	copy(b, a)
	fmt.Printf("a 长度%v 容量%v 值%v \n", len(a), cap(a), a)
	fmt.Printf("b 长度%v 容量%v 值%v \n", len(b), cap(b), b)
	b[0] = 789
	fmt.Printf("赋值后a 长度%v 容量%v 值%v \n", len(a), cap(a), a)
	fmt.Printf("赋值后b 长度%v 容量%v 值%v \n", len(b), cap(b), b)
}

func sortSlice() {
	aSlice := []int{3, 1, 5, 2, 8}
	sort.Ints(aSlice)
	fmt.Println(aSlice)
	bSlice := []string{"a", "c", "d", "b"}
	sort.Strings(bSlice)
	fmt.Println(bSlice)
}

func reserveSortSlice() {
	aSlice := []int{3, 1, 5, 2, 8}
	sort.Sort(sort.Reverse(sort.IntSlice(aSlice)))
	fmt.Println(aSlice)
	bSlice := []string{"a", "c", "d", "b"}
	sort.Sort(sort.Reverse(sort.StringSlice(bSlice)))
	fmt.Println(bSlice)
}

func main() {
	//createSlice()
	//showCap()
	//createArraySlice()
	//makeCreateSlice()
	//appendMethod()
	//appendManyMethod()
	//delMethod()
	//mergeSlice()
	//copySlice()
	//deepcopySlice()
	//sortSlice()
	reserveSortSlice()

}
