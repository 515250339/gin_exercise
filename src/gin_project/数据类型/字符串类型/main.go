package main

import (
	"fmt"
	"strings"
)

func createStr() {
	// 创建字符串
	fmt.Println("str := \"c:\\Code\\demo\\go.exe\"") // str := "c:\Code\demo\go.exe"
	str := "c:\\Code\\demo\\go.exe"
	fmt.Println("str := ", str)
}

func createManyLineStr() {
	// 创建的多行字符串
	sManyLine := `第一行
	第二行
	第三行`
	fmt.Println(sManyLine)

}

// byte和rune
func byteAndRune() {
	s := "日本垃垃"
	fmt.Println(s)
	sRune := []rune(s)
	fmt.Println("美国" + string(sRune[2:]))
}

// 查看字符串长度
func lenStr(a *string) {
	fmt.Println(len(*a))
}

// 拼接字符串
func strJoin(a, b *string) {
	fmt.Println(*a + *b)
}

// 切分字符串
func strSplit(a *string) {
	bArr := strings.Split(*a, "/")
	fmt.Println(bArr)
	fmt.Println(bArr[0])
}

// 判断字符串首尾字符
func strHasPrefix(a *string) {
	// 是否以test开头
	if strings.HasPrefix(*a, "test") {
		fmt.Println("是test开头")
	} else if strings.HasSuffix(*a, "end") { //是否以en结尾
		fmt.Println("是end结尾")
	}
}

func indexStr(a *string) {
	index := strings.Index(*a, "w")
	fmt.Println(index)
}

// 拼接字符串
func strJoin2(a *string) {
	aArr := strings.Split(*a, " ")
	fmt.Println(aArr)
	aStr := strings.Join(aArr, "-")
	fmt.Println(aStr)
}

func strASCII() {
	a := 'a'
	name := "张三"
	// 当我们直接输出 byte（字符）的时候输出的是这个字符对应的码
	fmt.Println(a) // 97 这里输出的是 a 字符串的 ASCII值
	fmt.Println(name)
	// 如果我们要输出这个字符，需要格式化输出
	fmt.Printf("%c\n", a)
	// 或者声明并强转为str
	var b = string('a')
	fmt.Println(b)
}

func main() {
	// rune方法
	//byteAndRune()

	// 拼接字符串
	//var a = "hi  "
	//var b = "lao 8"
	//strJoin(&a, &b)

	// 切分字符串
	//a := "123/456/789"
	//strSplit(&a)

	// 判断字符串首尾字符
	//a := "test_a"
	//strHasPrefix(&a)
	//a = "a_end"
	//strHasPrefix(&a)

	// 判断字符串出现的位置
	//a := "teswt"
	//indexStr(&a)

	// join拼接字符串
	//var b = "lao 8"
	//strJoin2(&b)

	// 单引号
	strASCII()
}
