package main

import "fmt"

// User 接口是一个规范
type User interface {
	start()
	end()
}

// Phone 如果接口里面有方法的话，必要要通过结构体或者通过自定义类型实现这个接口
type Phone struct {
	Name string
}

// 手机实现开机方法
func (p Phone) start() {
	fmt.Println(p.Name, "开机")
}

// 手机实现关机方法
func (p Phone) end() {
	fmt.Println(p.Name, "关机")
}

func main() {
	//p := Phone{Name: "Iphone"}
	//var p1 User // golang中接口就是一个数据类型
	//p1 = p      // 表示获取手机数据
	//p1.start()
	//p1.end()

	i := Phone{Name: "苹果"}
	var p1 User = i //苹果 实现了 开关机 接口 p1 是 Phone 类型
	p1.start()
	p1.end()

	m := &Phone{Name: "小米"}
	var p2 User = m //小米 实现了 开关机 接口 p2 是 *Phone 类型
	p2.start()
	p2.end()

}
