package main

import "fmt"

//空接口作为函数的参数
func show(a interface{}) {
	fmt.Printf("type：%T  value:%[1]v \n", a)
}

func main() {
	//show(20)
	//show("hello word")
	//slice := []int{1, 2, 3, 4}
	//show(slice)

	// 空切片作为 map 值
	//var slice = []interface{}{1, 2, 3, 4}
	//fmt.Printf("type：%T  value:%[1]v \n", slice)

	// 空接口作为 map 值
	var aMap = make(map[string]interface{})
	aMap["age"] = 123
	aMap["name"] = "张三"
	fmt.Printf("type：%T  value:%[1]v \n", aMap)
}
