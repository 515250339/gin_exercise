package main

import "fmt"

func main() {
	// 修改字符串，需要先将其转换成[]rune 或[]byte，完成后再转换为 string
	s1 := "big"
	byteS1 := []byte(s1)
	byteS1[0] = 'B'
	fmt.Println(string(byteS1))

	s2 := "Big"
	byteS2 := []rune(s2)
	byteS2[0] = 'b'
	fmt.Println(string(byteS2))
}
