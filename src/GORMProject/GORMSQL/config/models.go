package config

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID           uint
	Name         string
	Email        Email
	Age          uint8
	Birthday     time.Time
	MemberNumber sql.NullString
	ActivedAt    sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// TableName 定义结构体操作的数据库表
func (User) TableName() string {
	return "user"
}

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

type UserInfo struct {
	gorm.Model
	Name         string
	CreditCardId uint
	CreditCard   CreditCard `gorm:"foreignKey:UserID;references:CreditCardId;"`
}

type DefaultTest1 struct {
	ID   int64
	Name string `gorm:"default:gorm"`
	Age  int64  `gorm:"default:18"`
}

type DefaultTest2 struct {
	gorm.Model
	Name   string
	Age    *int         `gorm:"default:18"`
	Active sql.NullBool `gorm:"default:true"`
}

//type DefaultTest3 struct {
//	ID        string `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
//	FirstName string
//	LastName  string
//	Age       uint8
//	FullName  string `gorm:"type:GENERATED ALWAYS AS (concat(first_name,' ',last_name));default:(-);"`
//}
