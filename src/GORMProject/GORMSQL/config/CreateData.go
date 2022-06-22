package config

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

// CreateUser1 普通创建
func CreateUser1(db *gorm.DB) {
	user := User{Name: "z2s", Age: 18, Birthday: time.Now()}
	//result := db.Create(&user)       // 通过数据的指针来创建
	result := db.Session(&gorm.Session{SkipHooks: true}).Create(&user) // 通过数据的指针来创建
	fmt.Println(result.Error)                                          // 返回 error
	fmt.Println(result.RowsAffected)                                   // 返回插入记录的数量
}

// CreateUser2 指定字段创建
func CreateUser2(db *gorm.DB) {
	user := User{Name: "ls", Age: 22, CreatedAt: time.Now()}
	db.Select("Name", "Age", "CreatedAt").Create(&user)

	db.Omit("Name", "Age", "CreatedAt").Create(&user)
}

// CreateUser3 批量创建
func CreateUser3(db *gorm.DB) {
	users := []User{{Name: "qq"}, {Name: "ee"}}
	db.Create(&users)

	for _, user := range users {
		fmt.Println(user.ID)
	}
}

// CreateUser4 批量创建，指定批量大小
func CreateUser4(db *gorm.DB) {
	users := []User{{Name: "qq3"}, {Name: "ee3"}}
	db.CreateInBatches(users, 1)

	for _, user := range users {
		fmt.Println(user.ID)
	}
}

//func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
//	if u.Name != "zs" {
//		return errors.New("no zs")
//	}
//	return
//}

func CreateUser5(db *gorm.DB) {
	db.Model(&User{}).Create(map[string]interface{}{
		"Name": "aa", "Age": 19,
	})

	// batch insert from `[]map[string]interface{}{}`
	db.Model(&User{}).Create([]map[string]interface{}{
		{"Name": "aa2", "Age": 20},
		{"Name": "aa3", "Age": 21},
	})
}

func CreateUser6(db *gorm.DB) {
	// 通过 map 创建记录
	// clause.Expr{SQL: "类型(?)", Vars: []interface{}{"类型对应数据"}},
	db.Model(User{}).Create(map[string]interface{}{
		"Name": clause.Expr{SQL: "(?)", Vars: []interface{}{
			"zz",
		}},
	})

}

// 通过自定义累心创建
type Email struct {
	e string
}

func (em *Email) Scan(v interface{}) error {
	// Scan a value into struct from database driver
	return errors.New("no")
}

func (em Email) GormDataType() string {
	return "geometry"
}

func (em Email) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL: "(?)", Vars: []interface{}{fmt.Sprintf("%s", em.e)},
	}
}

func CreateUser7(db *gorm.DB) {
	db.Create(&User{
		Name: "test",
		Email: Email{
			e: "515@qq.com",
		},
	})
}

// CreateUser8 关联创建
func CreateUser8(db *gorm.DB) {
	//db.Create(&UserInfo{
	//	Name:       "zs",
	//	CreditCard: CreditCard{Number: "111111111111"},
	//})
	userinfo := UserInfo{Name: "test"}
	//db.Omit("CreditCard").Create(&userinfo)
	db.Omit(clause.Associations).Create(&userinfo)
}

// CreateUser9 创建带有默认值的表
func CreateUser9(db *gorm.DB) {
	//db.Create(&DefaultTest1{
	//	ID: 1,
	//})
	//
	//db.Create(&DefaultTest2{
	//	Name: "zs",
	//})
	//db.Create(&DefaultTest3{
	//	FirstName: "zhang",
	//	LastName:  "san",
	//	Age:       23,
	//})

}

func UpsertFunc(db *gorm.DB) {
	//db.Create(&DefaultTest1{
	//	ID: 1,
	//})
	//上面会抛出异常 Error 1062: Duplicate entry '1' for key 'default_test1.PRIMARY'

	// 有冲突时什么都不做
	//db.Clauses(clause.OnConflict{DoNothing: true}).Create(&DefaultTest1{
	//	ID: 1,
	//})

	// 当有冲突时，更新指定列为默认值
	//db.Clauses(clause.OnConflict{
	//	Columns: []clause.Column{
	//		{Name: "id"},
	//	},
	//	DoUpdates: clause.Assignments(map[string]interface{}{"id": 5}),
	//}).Create(&DefaultTest1{
	//	ID: 1,
	//})

	// 当有冲突时，更新指定列为新值
	db.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "id"},
		},
		DoUpdates: clause.AssignmentColumns([]string{"id", "age"}),
	}).Create(&DefaultTest1{
		ID:   1,
		Age:  20,
		Name: "zs",
	})

	db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&DefaultTest1{
		ID:  1,
		Age: 21,
	})
}
