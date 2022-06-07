package main

import "fmt"

// Age 我的年龄
type Age interface {
	MyAge()
}

// Sex 我的性别
type Sex interface {
	MySex()
}

// SetAgeAndSex 修改我的信息
type SetAgeAndSex interface {
	SetMyInfo(*string, *int)
}

// My 我的信息
type My struct {
	Sex string
	Age int
}

// SetMyInfo 修改我的信息
func (m *My) SetMyInfo(s *string, a *int) {
	m.Sex = *s
	m.Age = *a
	fmt.Println("修改信息")
}

// MySex 我的性别
func (m My) MySex() {
	fmt.Println("我的性别为", m.Sex)
}

// MyAge 我的年龄
func (m *My) MyAge() {
	fmt.Println("我", m.Age, "岁")
}

func main() {
	// 初始化我的信息
	my := &My{
		Sex: "男",
		Age: 998,
	}
	var age Age = my
	var sex Sex = my
	// 输出我的信息
	age.MyAge()
	sex.MySex()
	// 获取我的信息
	var setInfo SetAgeAndSex = my
	s := "娚"
	a := 18
	// 修改我的信息
	setInfo.SetMyInfo(&s, &a)
	age.MyAge()
	sex.MySex()

}
