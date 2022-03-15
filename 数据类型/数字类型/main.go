package main

import (
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
)

// 查看占用内存 参数aa 为int类型 从地址取值
func sizeOf(aa *int) {
	fmt.Println(*aa)
	// 输出类型为 %T
	fmt.Printf("%T\n", *aa)
	// 查看占用内存
	fmt.Println(unsafe.Sizeof(*aa))
}

// 类型转换 参数aa 为int类型 从地址取值
func transitionInt32(aa *int) {
	fmt.Println(*aa)
	// 输出类型为 %T
	fmt.Printf("%T\n", *aa)
	// 查看占用内存
	fmt.Println(unsafe.Sizeof(*aa))
	// 类型转换
	dd := int32(*aa)
	fmt.Println(dd)
	// 输出类型为 %T
	fmt.Printf("%T\n", dd)
	// 查看占用内存
	fmt.Println(unsafe.Sizeof(dd))
}

// 查看float类型 参数aa float64 从地址取值
func floatTest(aa *float64) {
	fmt.Println(*aa)
	// 输出类型为 %T
	fmt.Printf("%T\n", *aa)
	// 查看占用内存
	fmt.Println(unsafe.Sizeof(*aa))
}

// 查看数据类型
func showType() {
	a := 10
	fmt.Println(reflect.TypeOf(a))
}

// int 常用转换
func intTransition() {
	aStr := "123465"
	// 转为int32
	intV, _ := strconv.Atoi(aStr)
	fmt.Printf("intV:%v--type:%T\n", intV, intV)

	// 转为int64 base:10进制 bitSize:预期数值的bit大小，用于数值上限限制，最终返回的还是int64类型
	int64V, _ := strconv.ParseInt(aStr, 10, 64)
	fmt.Printf("intV:%v--type:%T\n", int64V, int64V)

	// int转string
	aInt := 123
	strS := strconv.Itoa(aInt)
	fmt.Printf("intV:%v--type:%T\n", strS, strS)
	// int64转string
	var aInt64 int64
	aInt64 = 123
	str64S := strconv.FormatInt(aInt64, 10)
	fmt.Printf("intV:%v--type:%T\n", str64S, str64S)
}

func main() {
	intTransition()
}
