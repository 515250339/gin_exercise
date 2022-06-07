package main

import "fmt"

// break 跳出当前循环，可⽤于 for、switch、select。
func breakTest() {
	fmt.Println("")
	fmt.Println("---------break-----------")

	for i := 1; 1 < 10; i++ {
		if i == 6 {
			break
		}
		fmt.Println(i)
	}
}

// continue 跳过本次循环，只能用于 for。
func continueTest() {
	fmt.Println("")
	fmt.Println("---------continue-----------")

	for i := 1; i < 10; i++ {
		if i == 7 {
			continue
		}
		fmt.Println(i)
	}
}

// goto 改变函数内代码执行顺序，不能跨函数使用。
func gotoTest() {
	fmt.Println("")
	fmt.Println("---------goto-----------")

	for i := 1; i < 10; i++ {
		if i == 6 {
			goto END
		}
		fmt.Println(i)
	}

END:
	fmt.Println("end")
}

func main() {
	//for m := 1; m < 10; m++ {
	//	/*    fmt.Printf("第%d次：\n",m) */
	//	for n := 1; n <= m; n++ {
	//		fmt.Printf("%dx%d=%d ", n, m, m*n)
	//	}
	//	fmt.Println("")
	//}
	gotoTest()
}
