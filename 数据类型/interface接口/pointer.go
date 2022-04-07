package main

import "fmt"

// Usb 接口定义
type Usb interface {
	start()
	end()
}

// MyPhone 结构体
type MyPhone struct {
	Name string
}

func (p *MyPhone) start() {
	fmt.Println(p.Name, "开启了Usb接口")
}

func (p *MyPhone) end() {
	fmt.Println(p.Name, "结束了Usb接口")
}

func main() {
	// 错误写法
	//i := MyPhone{Name: "华为手机"}
	//var p Usb = i
	//p.start()
	//p.end()

	// 正确写法
	i := MyPhone{Name: "华为手机"}
	var p Usb = &i
	p.start()
	p.end()
}
