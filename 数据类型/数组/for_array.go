package main

import "fmt"

func sumArr(a [3]int) {
	sum := 0
	for _, i := range a {
		sum += i
	}
	fmt.Printf("sum: %v, mean: %v", sum, float64(sum)/float64(len(a)))
}

func maxArr() {
	a := [4]int{1, 2, 4, 3}
	var max int = a[0]
	for _, i := range a {
		if i > max {
			max = i
		}
	}
	fmt.Println(max)
}

func main() {
	// 普通遍历
	//var a = [...]int{998, 2, 3}
	//for x := 0; x < len(a); x++ {
	//	fmt.Println(a[x])
	//}

	//// k,v遍历
	//for k, v := range a {
	//	fmt.Println(k, v)
	//}

	//sumArr(a)
	maxArr()
}
