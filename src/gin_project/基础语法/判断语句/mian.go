package main

import "fmt"

// switch 相当于 if elif
func switchTest() {
	fmt.Println("")
	fmt.Println("---------switch-----------")

	for i := 1; i <= 7; i++ {

		fmt.Printf("当 i = %d 时：\n", i)

		switch i {
		case 1:
			fmt.Println("输出 i=", i)
		case 2:
			fmt.Println("输出 i=", i)
		case 3:
			fmt.Println("输出 i=", i)
		case 4, 5, 6:
			fmt.Println("输出 i=", i)
		default:
			fmt.Println("输出 i", "默认")
		}
	}
}

func main() {
	switchTest()
}
