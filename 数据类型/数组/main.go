package main

import "fmt"

// 创建数组并赋值
func createArr() {
	// 创建数组 5 个元素 int 类型
	var a [5]int
	// 创建数组 3 个元素 int 类型
	var b [3]int
	fmt.Printf("类型：%T value：%[1]v \n", a)
	fmt.Printf("类型：%T value：%[1]v \n", b)
	// 进行下标取值并赋值 0 开始
	a[0] = 11
	a[3] = 33
	b[0] = 22
	b[2] = 99
	fmt.Printf("类型：%T value：%[1]v \n", a)
	fmt.Printf("类型：%T value：%[1]v \n", b)
}

// 数组是值类型
func updateArr(x [3]int) {
	x[2] = 100
	fmt.Println("【updateArr】-- 修改值为副本不影响本身的值", x)
}

// 数组是值类型
func updateArr2(x *[3]int) {
	x[2] = 100
	fmt.Println("【updateArr】-- 修改值为本身", *x)
}

func createCustomArr() {
	var a [3]int
	fmt.Println(a)

	var b = [3]int{1, 2}
	fmt.Println(b)
}

func createCustomArr2() {
	// 不定长Arr
	var c = []string{"张三", "李四"}
	fmt.Printf("类型：%T  value：%[1]v \n", c)
	// 不定长Arr 两个参数长度就为2
	var d = [...]string{"张三", "dd"}
	fmt.Printf("类型：%T  value：%[1]v \n", d)
	// 不定长Arr 长度默认下标加1 4+1=5
	var e = [...]int{0: 3, 1: 5, 4: 6}
	fmt.Printf("类型：%T  value：%[1]v \n", e)
}

func createCustomArr3() {
	// 指定索引赋值 长度默认下标加1 4+1=5
	var e = [...]int{0: 3, 1: 5, 4: 6}
	fmt.Printf("类型：%T  value：%[1]v \n", e)
}

func manyArr() [3][2]string {
	var a = [3][2]string{
		{"张三", "男"},
		{"里斯", "女"},
		{"老六", "中"},
	}
	//fmt.Println(a)
	//fmt.Println(a[1][0])
	return a
}

func main() {
	//createArr()
	// 数组是值类型
	//var a [3]int
	//updateArr(a)
	//fmt.Println(a)
	//fmt.Println()
	// 修改本身需要传递内存地址
	//var b [3]int
	//updateArr2(&b)
	//fmt.Println(b)
	//createCustomArr3()
	a := manyArr()
	for _, i := range a {
		var name string
		var sex string
		for k, v := range i {
			if k == 0 {
				name = v
				continue
			} else if k == 1 {
				sex = v
			}
			fmt.Println("name:", name, "sex:", sex)

		}
	}
}
