package main

import "fmt"

// 常量声明
const boilingF = 212.0

//func main() {
//	var f = boilingF
//	var c = (f - 32) * 5 / 9
//	fmt.Printf("%g° or %g℃", f, c)
//}

func main() {
	// 变量声明
	const freezingF, boiling = 32.0, boilingF
	// 调用函数 fToC 计算 入参为 float64 类型，返回参数为 float64 类型
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF)) // "32°F = 0°C"
	fmt.Printf("%g°F = %g°C\n", boiling, fToC(boiling))     // "212°F = 100°C"
}

// 入参为 float64 类型，返回参数为 float64 类型
func fToC(f float64) float64 {
	// 计算并返回
	return (f - 32) * 5 / 9
}
