package main

import "fmt"

func forSlice(a *[]string) {
	for i := 0; i < len(*a); i++ {
		// *a[i] 语法错误 可以理解为 *(a[i]) 应当 (*a)[i]
		fmt.Println((*a)[i])
	}
}

func kv_forSlice(a *[]string) {
	for k, v := range *a {
		fmt.Println(k, v)
	}
}

func main() {
	var a = []string{"test1", "test2", "test3"}
	//forSlice(&a)

	kv_forSlice(&a)
}
