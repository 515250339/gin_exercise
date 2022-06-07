package models

type User struct {
	Id       int
	UserName string
	Age      int
	Email    string
	AddTime  int
}

// TableName 定义结构体操作的数据库表
func (User) TableName() string {
	return "user"
}
