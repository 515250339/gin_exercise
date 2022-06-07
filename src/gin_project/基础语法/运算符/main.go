package main

import (
	"fmt"
)

func operation() {
	// 运算符
	fmt.Println("10+3=", 10+3)
	fmt.Println("10-3=", 10-3)
	fmt.Println("10*3=", 10*3)
	//除法注意：如果运算的数都是整数，那么除后，去掉小数部分，保留整数部分
	fmt.Println("10/3=", 10/3)
	fmt.Println("10.0/3=", 10.0/3)
	// 取余注意 余数=被除数-（被除数/除数）*除数
	fmt.Println("10%3=", 10%3)
	fmt.Println("-10%3=", -10%3)
	fmt.Println("10%-3=", 10%-3)
	fmt.Println("-10%-3=", -10%-3)
}

func i_jia_jia() {
	// i ++
	var i int = 1
	i++
	fmt.Println(i)

}

func relationOperation() {
	n1 := 9
	n2 := 8
	fmt.Println(n1 == n2)
	fmt.Println(n1 != n2)
	fmt.Println(n1 > n2)
	fmt.Println(n1 >= n2)
	fmt.Println(n1 < n2)
	fmt.Println(n1 <= n2)
	flag := n1 > n1
	fmt.Println(flag)
}

func logicOperation() {
	var age int = 40
	if age > 30 && age < 50 {
		fmt.Println("ok1")
	}
	if age > 30 && age < 40 {
		fmt.Println("ok2")
	}
	if age > 30 || age < 50 {
		fmt.Println("ok3")
	}
	if age > 30 || age < 40 {
		fmt.Println("ok4")
	}
	if age > 30 {
		fmt.Println("ok5")
	}
	if !(age > 30) {
		fmt.Println("ok6")
	}
}

func assignmentOperation() {
	d := 8 + 2*8
	fmt.Println(d)

	x := 10
	x += 5
	fmt.Println("x += 5 的值:", x)
}

func main() {
	//operation()
	//i_jia_jia()
	//relationOperation()
	//logicOperation()
	assignmentOperation()
}
